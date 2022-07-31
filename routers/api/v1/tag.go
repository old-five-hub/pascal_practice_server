package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/e"
	"pascal_practice_server/service/tag_service"
)

func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	tags, err := tag_service.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tags,
	})
}
