package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载路由文件
	r := setupRouter()

	// 使用日志中间件,设置日志文件
	r.Use(gin.Logger())
	f, _ := os.Create("./log/server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 设置静态资源目录
	r.Static("/static", "./static")

	// 加载视图目录
	r.LoadHTMLGlob("./views/*")

	// 8080端口运行
	if err := r.Run(":8080"); err != nil {
		fmt.Println("[startup service failed] Error:\n\t", err)
	}
}
