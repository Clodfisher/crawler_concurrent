package engine

import (
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

		parserResult, err := worker(&r)
		if err != nil {
			continue
		}
		requestSlice = append(requestSlice, parserResult.RequestSlice...)

		for _, item := range parserResult.ItemSlice {
			log.Printf("Got item %v", item)
		}
	}

}
