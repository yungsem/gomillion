package main

import (
	"fmt"
	"github.com/yungsem/gomillion/pkg/router"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic value: %v\n", r)
		}
	}()

	// 初始化路由
	r := router.Init()

	// 运行服务
	err := r.Run(":9700")
	if err != nil {
		panic(err)
	}
}
