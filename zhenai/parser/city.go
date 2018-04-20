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
		name := string(m[2])
		result.ItemSlice = append(result.ItemSlice, "User "+name)

		//用户的最新请求
		userRequest := engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParserResult {
				return ProfileParser(contents, name)
			},
		}
		result.RequestSlice = append(result.RequestSlice, userRequest)
	}

	return result
}
