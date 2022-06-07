package main

import(
    "reflect"

    "mime/multipart"

    "github.com/olivere/elastic/v7"

    "strings"

    "fmt"

    "errors"
)

const (
	POST_INDEX = "post"
)


type Post struct {
    Id      string `json:"id"` //raw string  
    User    string `json:"user"`
    Message string `json:"message"`
    Url     string `json:"url"`
    Type    string `json:"type"`
}

func searchPostByUser(user string) ([]Post, error) {
    /*

    type TermQuery struct {
	name            string
	value           interface{}
	boost           *float64
	caseInsensitive *bool
	queryName       string
    }

    func NewTermQuery(name string, value interface{}) *TermQuery {
        return &TermQuery{name: name, value: value}
    }
    */
    //interface{} type用于不知道什么类型的input
    //以下用于term搜索
    termQuery := elastic.NewTermQuery("user", user)

    searchResult, err := readFormES(termQuery, POST_INDEX)
    if err != nil {
		return nil,err
	}

    var posts []Post
    var post_type Post
    for _, item := range searchResult.Each(reflect.TypeOf(post_type)) {
		t := item.(Post)
        posts = append(posts, t)

	}

    return posts, nil
}

func searchPostByKeywords(keywords string) ([]Post, error) {
    /*
    func (q *MatchQuery) Operator(operator string) *MatchQuery {
	    q.operator = operator
	    return q
    }
    
    (q *MatchQuery) 就好像是 self.operator = operator, 就一constructor

    */

    matchQuery := elastic.NewMatchQuery("message", keywords)
    matchQuery.Operator("AND")

    if keywords == "" {
        matchQuery.ZeroTermsQuery("all")
    }


    searchResult, err := readFormES(matchQuery, POST_INDEX)
    if err != nil {
		return nil,err
	}
   
    var posts []Post
    var post_type Post
    for _, item := range searchResult.Each(reflect.TypeOf(post_type)) {
		t := item.(Post)
        posts = append(posts, t)

	}
    

    return posts, nil
}

func savePost(post *Post, file multipart.File) error {
    //past by reference can reduce operation complexity --> I directly use
    mediaLink, err := saveToGCS(file, post.Id)

    if err != nil {
        return err
    }

    post.Url = mediaLink
    return saveToES(post, POST_INDEX, post.Id)

}

func deletePost(post_id string, username string) (error, error) {

    fmt.Println(post_id, username)

    query := elastic.NewBoolQuery()
    query.Must(elastic.NewTermQuery("id", post_id)) //username0dgsdsd-asd-ddf
    query.Must(elastic.NewTermQuery("user", username))//username

    if !strings.Contains(post_id, username) {
        return errors.New("does not match"), errors.New("does not match")
    }

    err_ES := deleteFromES(query, POST_INDEX)
    err_GCS := deleteFromGCS(post_id)

    return err_ES, err_GCS

}



