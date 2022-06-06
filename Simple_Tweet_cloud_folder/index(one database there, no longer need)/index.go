package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

const (
	POST_INDEX = "post"
	USER_INDEX = "user"
	ES_URL = "http://10.128.0.3:9200" //internal IP在机器重启之后不会变，但？？？？
)



func main() {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("mingcheng", "123456789"))
	if err != nil {
		panic(err)
	}

	//检查和创建post database
	exists, err := client.IndexExists(POST_INDEX).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		//keyword with index true: 全文匹配 --> 将相关keyword抽取出来再提一个表用于快速匹配
		//text: 支持关键词搜索
		//keyword with index false: 没有快速全文匹配 
		mapping := `{
			"mappings": {
				"properties": {
					"id":       { "type": "keyword" },
					"user":     { "type": "keyword" },
					"message":  { "type": "text" },
					"url":      { "type": "keyword", "index": false },
					"type":     { "type": "keyword", "index": false }
				}
			}
		}`
		_, err := client.CreateIndex(POST_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

	//检查和创建user database
	exists, err = client.IndexExists(USER_INDEX).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		mapping := `{
			"mappings": {
				"properties": {
					"username": {"type": "keyword"},
					"password": {"type": "keyword"},
					"age":      {"type": "long", "index": false},
					"gender":   {"type": "keyword", "index": false}
				}
			}
		}`
		_, err = client.CreateIndex(USER_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Indexes are created.")
}