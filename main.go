package main

import (
	"github.com/Clodfisher/crawler_concurrent/engine"
	"github.com/Clodfisher/crawler_concurrent/persist"
	"github.com/Clodfisher/crawler_concurrent/scheduler"
	"github.com/Clodfisher/crawler_concurrent/zhenai/parser"
)

func main() {
	/*
		我们配置完itemsave后，配置Scheduler、WorkerCount、ItemChan。等全部配置好后运行
	*/
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		//SchedulerInterface: &scheduler.SimpleScheduler{},
		SchedulerInterface: &scheduler.QueuedScheduler{},
		WorkerCount:        100,
		ItemChan:           itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityListParser,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.CityParser,
	//})
}
