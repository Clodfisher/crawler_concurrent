package persist

import (
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Save: got item #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

//有两种方式，进行elasticsearch的存储，一种是http启用rest，一种是用elasticsearch客户端
func save(item interface{}) {
	//must sniff false in docker
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.79.133:9200/"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", resp)
}
