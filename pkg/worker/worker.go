package worker

import (
	"fmt"
	"github.com/yungsem/gomillion/pkg/param"
	"time"
)

type Worker struct {
	workQueue chan param.JobParam
}

func NewWorker() *Worker {
	return &Worker{workQueue: make(chan param.JobParam)}
}

func (w *Worker) Work(workQueueChan chan chan param.JobParam) {
	for {
		// 将自己的 workQueue 放入全局的 workQueueChan 中
		workQueueChan <- w.workQueue
		select {
		case payload := <-w.workQueue:
			// 模拟业务处理
			doBusiness(payload)
		}
	}
}

func doBusiness(payload param.JobParam) {
	fmt.Println(payload)
	time.Sleep(1 * time.Second)
}
