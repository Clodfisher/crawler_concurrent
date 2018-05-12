package persist

import (
	"context"
	"github.com/Clodfisher/crawler_concurrent/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	//must sniff false in docker
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.79.133:9200/"),
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Save: got item #%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver : error saving item %v : %v", item, err)
			}
		}
	}()
	return out, nil
}

//有两种方式，进行elasticsearch的存储，一种是http启用rest，一种是用elasticsearch客户端
func save(client *elastic.Client, index string, item engine.Item) error {
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
