package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 设置日志文件
	f, _ := os.Create("./log/server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 使用日志中间件
	r.Use(gin.Logger())

	// 设置静态文件夹
	r.Static("/static", "./static")

	// 加载路由
	//routers.LoadRouters(r)

	r.Run(":8080")
}
