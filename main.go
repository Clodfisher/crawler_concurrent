package main

import (
	"github.com/Clodfisher/crawler_concurrent/engine"
	"github.com/Clodfisher/crawler_concurrent/scheduler"
	"github.com/Clodfisher/crawler_concurrent/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		SchedulerInterface: &scheduler.SimpleScheduler{},
		WorkerCount:        100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityListParser,
	})
}
