package content

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/e"
	"pascal_practice_server/service/content_service"
)

type CreateQuestionForm struct {
	Name   string `form:"name" valid:"Required;MaxSize(100)"`
	TagIds []int  `form:"tagIds" valid:"Required"`
}

func CreateQuestion(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = CreateQuestionForm{}
	)

	c.BindJSON(&form)

	httpCode, errCode := app.Valid(&form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tags, err := content_service.GetTagByIds(form.TagIds)
	if len(tags) == 0 {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_QUESTION_FAIL, nil)
		return
	}

	err = content_service.CreateQuestion(form.Name, tags)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_QUESTION_FAIL, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type GetQuestionForm struct {
	TagIds []int `form:"tagIds"`
	Page   int   `form:"page" valid:"Required"`
	Limit  int   `form:"page" valid:"Required"`
}

func GetQuestionList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = GetQuestionForm{}
	)

	c.BindJSON(&form)

	httpCode, errCode := app.Valid(&form)

	if httpCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	result, err := content_service.GetQuestionList(form.TagIds, form.Page, form.Limit)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_QUESTION_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":    &result.List,
		"total":   &result.Total,
		"page":    form.Page,
		"limit":   form.Limit,
		"hasMore": &result.HasMore,
	})
}
