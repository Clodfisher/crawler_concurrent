package persist

import (
	"context"
	"encoding/json"
	"github.com/Clodfisher/crawler_concurrent/engine"
	"github.com/Clodfisher/crawler_concurrent/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:       "安静的雪",
			Gender:     "女",
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "人事/行政",
			Hukou:      "山东菏泽",
			Xingzuo:    "牡羊座",
			House:      "已购房",
			Car:        "未购车",
		},
	}

	// 通过Index,Type,Id三元组，反序列化成结构体
	// TODO:try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.79.133:9200/"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	//Save expected item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	//Fetch saved item
	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObject(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if expected != actual {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
