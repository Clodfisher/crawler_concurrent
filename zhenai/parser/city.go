package parser

import (
	"github.com/Clodfisher/crawler_concurrent/engine"
	"regexp"
)

const RegexpCity = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func CityParser(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(RegexpCity)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		//用户的名字
		result.ItemSlice = append(result.ItemSlice, "User "+string(m[2]))

		//用户的最新请求
		userRequest := engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		}
		result.RequestSlice = append(result.RequestSlice, userRequest)
	}

	return result
}
