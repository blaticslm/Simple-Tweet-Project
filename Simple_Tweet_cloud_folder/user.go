//jwt-go: generating
//go-jet-middleware: verifying

package main

import (
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

const (
	USER_INDEX = "user"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}

func checkUser(username, password string) (bool, error) {
	//bool: check existence
	//error: database error or not
	boolQuery := elastic.NewBoolQuery()

	/*
		func (q *BoolQuery) Must(queries ...Query) *BoolQuery {
			q.mustClauses = append(q.mustClauses, queries...)
			 return q
		}
	*/
	boolQuery.Must(elastic.NewTermQuery("username", username))
	boolQuery.Must(elastic.NewTermQuery("password", password))

	searchResult, err := readFormES(boolQuery, USER_INDEX)
	if err != nil {
		return false, err
	}

	var utype User
	for range searchResult.Each(reflect.TypeOf(utype)) {
		fmt.Printf("Login as %s\n", username)
		return true, nil //if you can get into this loop, you found the match user
	}

	return false, nil

}

func addUser(user *User) (bool, error) {
	//bool: add ok or not
	//error database error or not
	termQuery := elastic.NewTermQuery("username", user.Username)

	searchResult, err := readFormES(termQuery, USER_INDEX)
	if err != nil {
		return false, err
	}

	if searchResult.TotalHits() > 0 { //down below is also ok
		return false, nil
	}

	/*
		var utype User
		for _, item := range searchResult.Each(reflect.TypeOf(utype)){
			return false, nil //if you can get into this loop, you found the match user
		}
	*/

	err = saveToES(user, USER_INDEX, user.Username)
	if err != nil {
		return false, err
	}
	fmt.Printf("User is added: %s\n", user.Username)
	return true, nil

}
