package main

import (
	"pascal_practice_server/models"
	"pascal_practice_server/routers"
)

func init() {
	models.InitDb()
}

func main() {
	r := routers.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
