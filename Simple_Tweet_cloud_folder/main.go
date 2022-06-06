package main

import (
    "fmt"
    "log"
    "net/http"  //too weak

    "github.com/gorilla/mux"
    jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/form3tech-oss/jwt-go"

)

//why use this function as the junction for communication rather than putting all to front end?
//Answer: you can! but depending on aspects
//disadvantage: May slow down the webpage or device response; expode the database address
func main() {
    fmt.Println("started-service")

    //check token whether is matching
    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

    r := mux.NewRouter()
    //options是用来确跨域访问时对方域名是否支持本域名的请求
    //like pliot signal in information theory
    r.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST", "OPTIONS") 
    r.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET", "OPTIONS") 
    r.Handle("/post/{id}", jwtMiddleware.Handler(http.HandlerFunc(deleteHandler))).Methods("DELETE", "OPTIONS")
    r.Handle("/signup", http.HandlerFunc(signUpHandler)).Methods("POST", "OPTIONS") 
    r.Handle("/signin", http.HandlerFunc(signInHandler)).Methods("POST", "OPTIONS") 
    
    //http.HandleFunc("/upload", uploadHandler) //match url -- function
    log.Fatal(http.ListenAndServe(":8080", r)) //用于记录post通知

    //原理: 通过8080端口的http请求来搜索数据库
}
