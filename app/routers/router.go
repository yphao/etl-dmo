package routers

import (
	"etl-demo/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/6/26 18:07

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("views/*")

	//静态页面注册
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	//动态页面注册
	r.POST("/uploadFile", service.UpLoad)
	r.GET("/zipfileinfo", service.ZipFileInfo)
	r.GET("/extractfile", service.ExtractFileInfo)
	return r
}
