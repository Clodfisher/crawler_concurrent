package engine

import (
	"log"
)

type ConcurrentEngine struct {
	SchedulerInterface Scheduler
	WorkerCount        int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (ce *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan ParserResult)
	ce.SchedulerInterface.Run()

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(out, ce.SchedulerInterface)
	}

	for _, r := range seed {
		ce.SchedulerInterface.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.ItemSlice {
			log.Printf("Get item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, r := range result.RequestSlice {
			ce.SchedulerInterface.Submit(r)
		}
	}
}

func createWorker(out chan ParserResult, s Scheduler) {
	//每个worker都有一个自己的chanal，就自己告诉自己
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(&request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
