package scheduler

import "github.com/Clodfisher/crawler_concurrent/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (ss *SimpleScheduler) WorkerChan() chan engine.Request {
	return ss.workerChan
}

func (ss *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (ss *SimpleScheduler) Run() {
	ss.workerChan = make(chan engine.Request)
}

func (ss *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		ss.workerChan <- r
	}()
}
