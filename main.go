package main

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/15 20:46

import (
	"etl-demo/app/routers"
	"etl-demo/app/utils"
	"fmt"
)

func main() {
	config, err := utils.InitConfig()
	if err != nil {
		fmt.Println(err)
	}

	router := routers.SetupRouter()
	router.Run(config.Server.Ip + ":" + config.Server.Port)
}
