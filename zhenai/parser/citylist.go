package parser

import (
	"github.com/Clodfisher/crawler_concurrent/engine"
	"regexp"
)

const RegexpCityList = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func CityListParser(contents []byte, _ string) engine.ParserResult {
	re := regexp.MustCompile(RegexpCityList)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		//城市的名字,item用于存储，因此只保存有用的profile数据
		//result.ItemSlice = append(result.ItemSlice, "City "+string(m[2]))

		//城市的最新请求
		cityRequest := engine.Request{
			Url:        string(m[1]),
			ParserFunc: CityParser,
		}
		result.RequestSlice = append(result.RequestSlice, cityRequest)

	}

	return result
}
