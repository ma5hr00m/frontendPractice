package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 路由函数

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", helloHandler)
	return r
}

// 视图函数

func helloHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
