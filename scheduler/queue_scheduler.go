package scheduler

import "github.com/Clodfisher/crawler_concurrent/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	//每个worker有个单独的chanal接口
	workerChan chan chan engine.Request
}

func (qs *QueuedScheduler) Submit(r engine.Request) {
	qs.requestChan <- r
}

func (qs *QueuedScheduler) WorkerReady(w chan engine.Request) {
	qs.workerChan <- w
}

func (qs *QueuedScheduler) ConfigMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

/*
两个独立的事情同时去收，用select
同时收到一个request让其排队，收到一个worker也让其排队
nil同时也不会诶select到
*/
func (qs *QueuedScheduler) Run() {
	qs.requestChan = make(chan engine.Request)
	qs.workerChan = make(chan chan engine.Request)
	go func() {
		var queueRequest []engine.Request
		var queueWorker []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(queueRequest) > 0 && len(queueWorker) > 0 {
				activeRequest = queueRequest[0]
				activeWorker = queueWorker[0]
			}
			select {
			case r := <-qs.requestChan:
				queueRequest = append(queueRequest, r)
			case w := <-qs.workerChan:
				queueWorker = append(queueWorker, w)
			case activeWorker <- activeRequest:
				queueRequest = queueRequest[1:]
				queueWorker = queueWorker[1:]
			}
		}
	}()
}
