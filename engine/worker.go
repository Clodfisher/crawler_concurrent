package engine

import (
	"github.com/Clodfisher/crawler_concurrent/fetcher"
	"log"
)

func worker(r *Request) (ParserResult, error) {
	//将请求递交给fetch获取网页内容text
	log.Printf("Fetching Url: %s", r.Url)
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	//将网页内容text交给解析器处理
	parserResult := r.ParserFunc(contents, r.Url)

	return parserResult, nil
}
