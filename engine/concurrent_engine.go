package engine

import (
	"log"
)

type ConcurrentEngine struct {
	SchedulerInterface Scheduler
	WorkerCount        int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	//我有个worker请问给我那个chan
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (ce *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan ParserResult)
	ce.SchedulerInterface.Run()

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.SchedulerInterface.WorkerChan(), out, ce.SchedulerInterface)
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

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	//每个worker都有一个自己的chanal，就自己告诉自己
	//in := make(chan Request)
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(&request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
