package service

import (
	"etl-demo/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/27 15:30

func ZipFileInfo(c *gin.Context) {

	//ufi  := models.UpLoadFile{}
	//var _ufi =[] models.UpLoadFile
	ufi, err := models.QueryAllZipInfo()
	if err != nil {

		panic(err)
	}

	//fmt.Println(reflect.TypeOf(ufi))

	c.JSON(http.StatusOK, &ufi)
	//c.HTML(http.StatusOK, "allzipfileinfo.html", gin.H{
	//	"fileinfo": &ufi,
	//})
}
