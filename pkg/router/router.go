package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yungsem/gomillion/pkg/handler"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// 分配 ID
	r.POST("/sendData", handler.SendData)

	return r
}
