package content

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/e"
	"pascal_practice_server/service/content_service"
)

func GetAllTags(c *gin.Context) {
	appG := app.Gin{C: c}
	tags, err := content_service.GetAllTags()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
	}
	count := len(tags)
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

type CreateTagForm struct {
	Name string `form:"name" valid:"MaxSize(100)"`
	Hot  int    `form:"hot" valid:"Range(0, 1)"`
	Icon string `form:"icon" valid:"Required"`
}

func CreateTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = CreateTagForm{}
	)

	c.BindJSON(&form)

	httpCode, errCode := app.Valid(form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	exists, err := content_service.ExistTagByName(form.Name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}
	tag := content_service.Tag{
		Name: form.Name,
		Hot:  form.Hot,
	}

	err = content_service.CreateTag(&tag)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

type UpdateTagForm struct {
	Name string `form:"name" valid:"required;MaxSize(100)"`
	Hot  int    `form:"hot" valid:"Range(0,1)"`
	Icon string `form:"icon" valid:"Required"`
	Id   int    `form:"id"`
}

func UpdateTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form UpdateTagForm
	)

	c.BindJSON(&form)

	httpCode, errCode := app.Valid(form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tag := content_service.Tag{
		Id:   form.Id,
		Name: form.Name,
		Hot:  form.Hot,
		Icon: form.Icon,
	}

	exists, err := content_service.ExistTagById(tag.Id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}
	err = content_service.UpdateTag(&tag)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	tag := content_service.Tag{Id: id}
	exists, err := content_service.ExistTagById(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err := content_service.DeleteTag(tag.Id); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
