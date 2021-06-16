package main

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/15 20:46

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Set a lower memory limit for multipart forms, default is 32MB
	r.MaxMultipartMemory = 8 << 20 // 8MB
	r.Static("/", "./views")
	r.POST("/upload", func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		filename := filepath.Base(file.Filename)
		upLoadFileName := fmt.Sprintf(`./upload/` + filename)
		if err := c.SaveUploadedFile(file, upLoadFileName); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.String(http.StatusOK,
			fmt.Sprintf("File %s uploaded successfully.", filename))
	})
	r.Run(":8088")
}
