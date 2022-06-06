 package main

 import(
	 "context"

	 "github.com/olivere/elastic/v7"
 )

 const(
	ES_URL = "http://10.128.0.3:9200"
 )

 func readFormES(termQuery elastic.Query, index string)(*elastic.SearchResult, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("mingcheng", "123456789"))
	if err != nil {
		return nil,err
	}

	//example-test.go
	searchResult, err := client.Search().
		Index(index).        // search in index "twitter"
		Query(termQuery).        // specify the query
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		return nil,err
	}

	return searchResult, nil

 }
//input_data is not comfirmed: it can be both post and user 
 func saveToES(input_data interface{}, index string, primaryKey string) error{
    client, err := elastic.NewClient(
        elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("mingcheng", "123456789"))
    if err != nil {
        return err
    }

	//i can be anything, as far as it BodyJson(i) can be successful 
    _, err = client.Index().
        Index(index).
        Id(primaryKey).
        BodyJson(input_data).
        Do(context.Background())
    return err
}

func deleteFromES(query elastic.Query, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("mingcheng", "123456789"))
	if err != nil {
		return err
	}

	//example-test.go
	_, err = client.DeleteByQuery().
		Index(index).        // search in index "twitter"
		Query(query).        // specify the query
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
		
	return err
}
