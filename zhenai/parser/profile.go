package parser

import (
	"github.com/Clodfisher/crawler_concurrent/engine"
	"github.com/Clodfisher/crawler_concurrent/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)</span></td>`)

var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(` <td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var xingzuo = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)

var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ProfileParser(contents []byte, url string, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Age = extractInt(contents, ageRe)
	profile.Height = extractInt(contents, heightRe)
	profile.Weight = extractInt(contents, weightRe)

	profile.Marriage = extractString(contents, marriageRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Education = extractString(contents, educationRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	profile.Xingzuo = extractString(contents, xingzuo)

	result := engine.ParserResult{
		ItemSlice: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		result.RequestSlice = append(result.RequestSlice,
			engine.Request{
				Url:        url,
				ParserFunc: ParserProfile(name),
			})
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	m := re.FindSubmatch(contents)
	if len(m) >= 2 {
		return string(m[1])
	} else {
		return ""
	}
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	i, err := strconv.Atoi(extractString(contents, re))
	if err == nil {
		return i
	}

	return 0
}

func ParserProfile(name string) engine.ParserFuncType {
	return func(c []byte, url string) engine.ParserResult {
		return ProfileParser(c, url, name)
	}
}
