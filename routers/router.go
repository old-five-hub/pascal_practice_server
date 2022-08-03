package routers

import (
	"github.com/gin-gonic/gin"
	"pascal_practice_server/middleware/jwt"
	v1 "pascal_practice_server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	appV1 := r.Group("/app/v1")
	appV1.Use(jwt.JWT())
	{
		appV1.POST("/tags", v1.GetTags)
	}

	return r
}
