package main

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/15 20:46

import (
	"etl-demo/app/models"
	"etl-demo/app/routers"
	"etl-demo/app/utils"
)

func main() {

	models.SetUp(false)

	config := utils.GetConfig()
	router := routers.SetupRouter()
	router.Run(config.Server.Ip + ":" + config.Server.Port)
}
