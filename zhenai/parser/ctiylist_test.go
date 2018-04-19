package parser

import (
	"io/ioutil"
	"testing"
)

func TestCityListParser(t *testing.T) {
	//	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	parserResult := CityListParser(contents)

	const resultSize = 470
	expectUrl := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(parserResult.RequestSlice) != resultSize {
		t.Errorf("result should have %d request, but had %d",
			resultSize, len(parserResult.RequestSlice))
	}
	for i, url := range expectUrl {
		if parserResult.RequestSlice[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s",
				i, url, parserResult.RequestSlice[i].Url)
		}

	}
}
