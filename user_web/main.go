package main

import (
	"fmt"
	"go.uber.org/zap"
	"project/catsshop_api/user_web/initialize"
)

func main() {
	port := 8021

	//1.初始化日志
	initialize.InitLogger()

	//2.初始化routers
	Router := initialize.Routers()

	zap.S().Infof("启动服务器，端口：%d", port)

	err := Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
