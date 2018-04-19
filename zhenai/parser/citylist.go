package parser

import (
	"github.com/Clodfisher/crawler_concurrent/engine"
	"regexp"
)

const RegexpCityList = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func CityListParser(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(RegexpCityList)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		//城市的名字
		result.ItemSlice = append(result.ItemSlice, string(m[2]))

		//城市的最新请求
		cityRequest := engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		}
		result.RequestSlice = append(result.RequestSlice, cityRequest)
	}

	return result
}
