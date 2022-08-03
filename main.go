package main

import (
	"pascal_practice_server/models"
	"pascal_practice_server/pkg/gredis"
	"pascal_practice_server/pkg/setting"
	"pascal_practice_server/pkg/utils"
	"pascal_practice_server/routers"
)

func init() {
	setting.SetUp()
	models.SetUp()
	gredis.SetUp()
	utils.SetUp()
}

func main() {
	r := routers.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
