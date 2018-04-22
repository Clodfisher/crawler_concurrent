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
}

func (ce *ConcurrentEngine) Run(seed ...Request) {

	in := make(chan Request)
	out := make(chan ParserResult)
	ce.SchedulerInterface.ConfigMasterWorkerChan(in)

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(in, out)
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

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(&request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
