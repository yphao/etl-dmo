package service

import (
	"etl-demo/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/27 22:03

func ExtractFileInfo(c *gin.Context) {

	efi, err := models.QueryALLExtractFileInfo()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, &efi)
}
