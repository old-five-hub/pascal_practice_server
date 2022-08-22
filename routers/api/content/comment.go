package content

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/e"
	"pascal_practice_server/service/content_service"
)

type GetCommentsForm struct {
	QuestionId int `form:"questionId"valid:"Required"`
}

func GetComments(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = GetCommentsForm{}
	)

	c.BindJSON(&form)

	httpCode, errCode := app.Valid(&form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	comments, err := content_service.GetComments(form.QuestionId)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"comments": comments,
	})
}

type CreateCommentForm struct {
	QuestionId int    `form:"questionId"valid:"Required"`
	Content    string `json:"content"valid:"Required"`
	ParentId   string `json:"parentId"`
}

func CreateComment(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = CreateCommentForm{}
	)

	c.BindJSON(&form)

	httpCode, errCode := app.Valid(&form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	err := content_service.CreateComment(map[string]interface{}{
		"questionId": form.QuestionId,
		"Content":    form.Content,
		"ParentId":   form.ParentId,
	})

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
