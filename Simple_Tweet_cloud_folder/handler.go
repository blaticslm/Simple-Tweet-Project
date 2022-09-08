package main

import(
	"encoding/json"
	"fmt"
	"net/http"
    "github.com/gorilla/mux"

    "path/filepath"
    "github.com/pborman/uuid"

    "regexp"
    "time"
    jwt "github.com/form3tech-oss/jwt-go"


)

var signingKey = []byte("secret") //本来就长这样，和其他类型的数据不一样

var (
    mediaTypes = map[string]string{
        ".jpeg": "image",
        ".jpg":  "image",
        ".gif":  "image",
        ".png":  "image",
        ".mov":  "video",
        ".mp4":  "video",
        ".avi":  "video",
        ".flv":  "video",
        ".wmv":  "video",
    }
)


//传request的地址的原因是要完全修改request的东西
func uploadHandler(w http.ResponseWriter, r *http.Request) {  
    //http.request is a struct, 所以支持指针
    //http.Responsewriter is interface，不支持pointer。意义: 
    
    fmt.Println("Received one", r.Method,"upload request")

    //具体source code: https://pkg.go.dev/net/http
    w.Header().Set("Access-Control-Allow-Origin", "*") //可以所有domain跨域访问
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //可以允许的http header，以及golang可以用这种方式来加到string array里， 详细看源代码

    //如果来的时option, 那么上面response的header信息就写入那两个允许的范围
    //
    if r.Method == "OPTIONS" {
        return
    }

    token := r.Context().Value("user")
    claims := token.(*jwt.Token).Claims //cast token to jwt.Token format, and then find out the payload
    username := claims.(jwt.MapClaims)["username"]

    //暂时就三个初始化
    //From now on, the user info will directly come from payload.
    p := Post{
        Id: username.(string) + uuid.New(), //universally unique identifier --- MOST IMPORTANT!!!!!!!!!!!!!!!!!
        User: username.(string), //func (r *Request) FormValue(key string) string
        Message: r.FormValue("message"),
    }

    //func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
    file, header, err := r.FormFile("media_file") //在postman中的其中一个key
    if err != nil {
        http.Error(w, "Media file is not avaliable", http.StatusBadRequest)
        return
    }

    suffix := filepath.Ext(header.Filename)
    t := mediaTypes[suffix]

    if t != "" {
        p.Type = t
    } else {
        p.Type = "unknown"
    }

    err = savePost(&p, file)
    if err != nil {
        http.Error(w, "Fail to save post to GCS or ES", http.StatusInternalServerError)
        return
    }

    fmt.Println("Post saved successfully")

    /*test code, there will be specific function later on
    //删除掉的原因是有data上传了，json格式就再也不太行了
    decoder := json.NewDecoder(r.Body) //解析得到request的body
    var p Post //p is struct
    if err := decoder.Decode(&p); err != nil { //然后讲解析得到的request和Post对应起来
        panic(err)
    }

    fmt.Fprintf(w, "Post received: %s\n", p.Message)
    */ 
}


func searchHandler(w http.ResponseWriter, r *http.Request) {

    //具体source code: https://pkg.go.dev/net/http
    w.Header().Set("Access-Control-Allow-Origin", "*") //可以所有domain跨域访问
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //可以允许的http header，以及golang可以用这种方式来加到string array里， 详细看源代码
    //告诉浏览器我即将返回的文件类型。upload不需要特定的类型，但是搜索返回的数据是要说清楚类型的
    w.Header().Set("Content-Type", "application/json")

    //如果来的search的option, 那么上面response的header信息就写入那两个允许的范围
    if r.Method == "OPTIONS" {
        fmt.Println("Received one search Option request")
        return
    }

    fmt.Println("Received one search request")
    //https://pkg.go.dev/net/url#Values.Get
    user := r.URL.Query().Get("user") //http://FRONT_URL?user=xxx
    keywords := r.URL.Query().Get("keywords")

    var posts []Post
    var err error
    if user != "" {
        posts, err = searchPostByUser(user)
    } else {
        posts, err = searchPostByKeywords(keywords)
    }

    //读取数据库出问题
    if err != nil{
        http.Error(w, "Failed to read post from the search", http.StatusInternalServerError)
        return
    }

    //Marshal returns the JSON encoding of posts. 
    js, err := json.Marshal(posts)

    //无法复写
    if err != nil{
        http.Error(w, "Failed to parse the post result to json", http.StatusInternalServerError)
        return
    }

    w.Write(js)

}

func signInHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*") 
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") 
    w.Header().Set("Content-Type", "text/plain") //token is a string text

    if r.Method == "OPTIONS" {
        fmt.Println("Received one login Option request")
        return
    }

    fmt.Println("Received one login request")
    //checkUser is in the database?
    decoder := json.NewDecoder(r.Body)
    var user User //p is struct
    if err := decoder.Decode(&user); err != nil { //然后讲解析得到的request和Post对应起来
        http.Error(w, "Failed to read user infomation from client", http.StatusBadRequest)
        return
    }

    check, err := checkUser(user.Username, user.Password) 
    if err != nil {
        http.Error(w, "Failed to read data from database", http.StatusInternalServerError)
        return
    }

    if !check {
        http.Error(w, "Wrong username/password or user doesn't exist", http.StatusUnauthorized)
        return
    }

    //create token


    token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(signingKey)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        fmt.Printf("Failed to generate token %v\n", err)
        return
    }
    
    //return token to client
    //checked in postman
    w.Write([]byte(tokenString))
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Received one signup request")

    w.Header().Set("Access-Control-Allow-Origin", "*") 
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") 
    //no need return any information

    if r.Method == "OPTIONS" {
        return
    }

    decoder := json.NewDecoder(r.Body)
    var user User //p is struct
    if err := decoder.Decode(&user); err != nil { //然后讲解析得到的request和Post对应起来
        http.Error(w, "Failed to read user information from client", http.StatusBadRequest)
        return
    }

    //Validate the format of the new  user data
    //can be done in front end
    //this regexp I going to verify
    if user.Username == "" || user.Password == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.Username) {
        http.Error(w, "Invalid username or password", http.StatusBadRequest)
        fmt.Printf("Invalid username or password\n")
        return
    }

    success, err := addUser(&user)
    if err != nil {
        http.Error(w, "Failed to sign up", http.StatusInternalServerError)
        return
    }

    if !success {
        http.Error(w, "Username exist", http.StatusBadRequest)
        return
    }

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*") 
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
    w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,DELETE") //browser may not recognized DELETE like post, get option

    //no need return any information

    if r.Method == "OPTIONS" {
        fmt.Println("Received one delete Option request")
        return
    }

    fmt.Println("Received one delete request")

    token := r.Context().Value("user")
    claims := token.(*jwt.Token).Claims //cast token to jwt.Token format, and then find out the payload
    username := claims.(jwt.MapClaims)["username"].(string)
    id := mux.Vars(r)["id"]


    err_ES, err_GCS := deletePost(id, username)

    if err_ES != nil && err_GCS == nil {
        http.Error(w, "Failed to delete post from Elasticsearch. Post may not exist.", http.StatusInternalServerError)
        fmt.Printf("Failed to delete post from Elasticsearch %v\n", err_ES)

    } else if err_ES == nil && err_GCS != nil {
        http.Error(w, "Failed to delete File from GCS. File may not exist.", http.StatusInternalServerError)
        fmt.Printf("Failed to delete post from GCS %v\n", err_GCS)

    } else if err_ES != nil && err_GCS != nil {
        http.Error(w, "Username does not match the post or both servers are gone, delete fail!", http.StatusInternalServerError)
        fmt.Printf("Id does not match to username or both servers are done!")

    }else {
        fmt.Println("Post is deleted totally successfully")
    }


}