package parser

import (
	"fmt"
	_ "github.com/Clodfisher/crawler_concurrent/fetcher"
	"io/ioutil"
	"testing"
)

func TestCityListParser(t *testing.T) {
	//	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", contents)
}
