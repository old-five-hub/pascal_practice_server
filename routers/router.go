package routers

import (
	"github.com/gin-gonic/gin"
	"pascal_practice_server/middleware/jwt"
	"pascal_practice_server/routers/api/account"
	"pascal_practice_server/routers/api/content"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/api/account/login", account.Login)

	appContent := r.Group("/api/content")
	appContent.Use(jwt.JWT())
	{
		appContent.POST("tag/list", content.GetAllTags)
		appContent.POST("tag/update", content.UpdateTag)
		appContent.POST("tag/create", content.CreateTag)
		appContent.POST("tag/delete", content.DeleteTag)
	}

	appAccount := r.Group("/api/account")
	appAccount.Use(jwt.JWT())
	{
		appAccount.POST("info", account.Info)
	}

	return r
}
