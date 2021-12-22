package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//html渲染

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*.html")
	r.GET("/test01", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "我是标题",
		})
	})
	r.GET("/test02", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "新闻页面",
		})
	})
	r.Run(":8089")
}
