package content

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pascal_practice_server/models"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/e"
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

	tags, err := models.GetTagByIds(form.TagIds)
	if len(tags) == 0 {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_QUESTION_FAIL, nil)
		return
	}

	err = models.CreateQuestion(form.Name, tags)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CREATE_QUESTION_FAIL, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
