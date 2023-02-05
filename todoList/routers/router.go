package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// ctrs "github.com/zachrey/blog/controllers"
)

// LoadRouters 初始化router
func LoadRouters(router *gin.Engine) {
	loadRouters(router)
}

func loadRouters(router *gin.Engine) {
	// 这里测试根路由
	router.GET("/", func(c *gin.Context) {
		post := models.GetPostByID(1)
		log.Println("@@ post", post)
		// 返回一个json格式的数据
		c.JSON(http.StatusOK, gin.H{
			"Status": 0,
			"data":   post,
		})
	})
	// 路由控制函数，我们全部放在controllers目录下
	//router.GET("/get-posts", ctrs.GetPosts)

	// ......  很多很多路由。。。
}
