package engine

import (
	"github.com/Clodfisher/crawler_concurrent/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (se SimpleEngine) Run(seed ...Request) {
	var requestSlice []Request
	for _, r := range seed {
		requestSlice = append(requestSlice, r)
	}

	for len(requestSlice) > 0 {

		//取得单个请求
		r := requestSlice[0]
		requestSlice = requestSlice[1:]

		parserResult, err := se.worker(&r)
		if err != nil {
			continue
		}
		requestSlice = append(requestSlice, parserResult.RequestSlice...)

		for _, item := range parserResult.ItemSlice {
			log.Printf("Got item %v", item)
		}
	}

}

func (se SimpleEngine) worker(r *Request) (ParserResult, error) {
	//将请求递交给fetch获取网页内容text
	log.Printf("Fetching Url: %s", r.Url)
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	//将网页内容text交给解析器处理
	parserResult := r.ParserFunc(contents)

	return parserResult, nil
}
