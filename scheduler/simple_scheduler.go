package scheduler

import "github.com/Clodfisher/crawler_concurrent/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (ss *SimpleScheduler) ConfigMasterWorkerChan(c chan engine.Request) {
	ss.workerChan = c
}

func (ss *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		ss.workerChan <- r
	}()
}
