package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
)

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func InitElastic() {
	url := elastic.SetURL("http://127.0.0.1:9200")
	sniffOpt := elastic.SetSniff(false)

	client, err := elastic.NewClient(url, sniffOpt)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	esVersion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esVersion)

	exists, err := client.IndexExists("twitter").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		mapping := `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"doc":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
                "retweets":{
                    "type":"long"
                },
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}
`
		createIndex, err := client.CreateIndex("twitter").Body(mapping).IncludeTypeName(true).Do(context.Background())
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			panic(createIndex)
		}
	}
	tweet1 := Tweet{User: "yu hao", Message: "Take Five", Retweets: 98}
	put1, err := client.Index().Index("twitter").Id("2").BodyJson(tweet1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	get1, err := client.Get().Index("twitter").Id("2").Do(context.Background())
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			panic(fmt.Sprintf("Document not found: %v", err))

		case elastic.IsTimeout(err):
			panic(fmt.Sprintf("Document is timeout: %v", err))

		case elastic.IsConnErr(err):
			panic(fmt.Sprintf("Connection problem: %v", err))

		default:
			// Some other kind of error
			panic(err)
		}
	}
	fmt.Printf("%s,%s,%s\n", get1.Id, get1.Index, get1.Source)

	fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)

	termQuery := elastic.NewTermQuery("user", "yu hao")
	searchResult, err := client.Search().
		Index("twitter").
		Query(termQuery).
		Sort("user", true).
		From(0).Size(10).
		Pretty(true).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(Tweet)
		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	}
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	if searchResult.Hits.TotalHits.Value > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits.Value)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Tweet
			err := json.Unmarshal(hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}

	// Update a tweet by the update API of Elasticsearch.
	// We just increment the number of retweets.
	script := elastic.NewScript("ctx._source.retweets += params.num").Param("num", 1)
	update, err := client.Update().Index("twitter").Id("1").
		Script(script).
		Upsert(map[string]interface{}{"retweets": 65}).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("New version of tweet %q is now %d", update.Id, update.Version)

	// ...
	//
	//// Delete an index.
	//deleteIndex, err := client.DeleteIndex("twitter").Do(context.Background())
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//if !deleteIndex.Acknowledged {
	//	// Not acknowledged
	//}
}
