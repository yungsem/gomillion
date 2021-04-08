package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yungsem/gomillion/pkg/dispatcher"
	"github.com/yungsem/gomillion/pkg/param"
	"github.com/yungsem/gomillion/pkg/worker"
	"github.com/yungsem/gotool/result"
	"net/http"
)

// 全局的 workQueueChan 的最大 size
// 一个 worKQueue 绑定一个 worker ，所以这个 size 也是 worker 的数量
const workQueueChanMaxSize = 200

// 创建全局的 workQueueChan
// dispatcher 从 workQueueChan 中阻塞等待，获取一个可用的 workQueue ，然后将从 jobQueue 中取出一个 payload 写入这个 workQueue 中
// worker 和 workQueue 是一一绑定的，worker 的 work 方法主要职责是不停地循环，将自己的 workQueue 写入全局的 workQueueChan 中，
// 并监听自己的 workQueue ，如果自己的 workQueue 中有可用的 payload ，就取出来作业务处理
var workQueueChan = make(chan chan param.JobParam, workQueueChanMaxSize)

// 创建全局的 dispatcher
var dspr = dispatcher.NewDispatcher()

func init() {
	// 创建 workQueueChanMaxSize 个 worker
	// 每个 worker 都单独开启一个 goroutine 来工作
	for i := 0; i < workQueueChanMaxSize; i++ {
		w := worker.NewWorker()
		go w.Work(workQueueChan)
	}

	// 启动 dispatcher 的 dispatch
	go dspr.DisPatch(workQueueChan)
}

// 发送数据
func SendData(c *gin.Context) {
	var eqpStatus param.EqpStatus
	if err := c.ShouldBindJSON(&eqpStatus); err != nil {
		c.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}

	dspr.Receive(&eqpStatus)

	c.JSON(http.StatusOK, result.Success("ok"))
}
