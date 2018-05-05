package persist

import (
	"github.com/Clodfisher/crawler_concurrent/model"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name:       "HiSiri",
		Gender:     "女",
		Age:        28,
		Height:     163,
		Weight:     100,
		Income:     "3001-5000元",
		Marriage:   "未婚",
		Education:  "大学本科",
		Occupation: "人事/行政",
		Hukou:      "内蒙古赤峰",
		Xingzuo:    "金牛座",
		House:      "自住",
		Car:        "未购车",
	}

	save(profile)
}
