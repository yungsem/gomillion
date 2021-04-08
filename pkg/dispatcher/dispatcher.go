package dispatcher

import "github.com/yungsem/gomillion/pkg/param"

const jobQueueMaxSize = 100

type Dispatcher struct {
	jobQueue chan param.JobParam
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{jobQueue: make(chan param.JobParam, jobQueueMaxSize)}
}

func (d *Dispatcher) Receive(payload param.JobParam) {
	d.jobQueue <- payload
}

func (d *Dispatcher) DisPatch(workQueueChan chan chan param.JobParam) {
	for {
		select {
		case payload := <-d.jobQueue:
			go func(payload param.JobParam) {
				workQueue := <-workQueueChan
				workQueue <- payload
			}(payload)
		}
	}
}
