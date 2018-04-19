package engine

import (
	"github.com/Clodfisher/crawler_concurrent/fetcher"
	"log"
)

func Run(seed ...Request) {
	var requestSlice []Request
	for _, r := range seed {
		requestSlice = append(requestSlice, r)
	}

	for len(requestSlice) > 0 {

		//取得单个请求
		r := requestSlice[0]
		requestSlice = requestSlice[1:]
		log.Printf("Fetching Url: %s", r.Url)

		//将请求递交给fetch获取网页内容text
		contents, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		//将网页内容text交给解析器处理
		parserResult := r.ParserFunc(contents)
		requestSlice = append(requestSlice, parserResult.RequestSlice...)

		for _, item := range parserResult.ItemSlice {
			log.Printf("Got item %v", item)
		}
	}

}
