package main

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/15 20:46

import (
	"etl-demo/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
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
		upLoadFileName := fmt.Sprintf(`./.upload_tmp/` + filename)
		if err != nil {
			panic(err)
		}
		if err := c.SaveUploadedFile(file, upLoadFileName); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		fileMd5, err := tools.CalcFileMd5(upLoadFileName)
		c.String(http.StatusOK,
			fmt.Sprintf("File %s uploaded successfully,file md5 %s", filename, fileMd5))
	})
	r.Run(":8088")
}
