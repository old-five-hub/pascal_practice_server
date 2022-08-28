package content

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pascal_practice_server/models"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/e"
	"pascal_practice_server/pkg/utils"
	"pascal_practice_server/service/content_service"
)

type GetCommentsForm struct {
	QuestionId int `form:"questionId"valid:"Required"`
}

type GetCommentResultItem struct {
	models.Comment
	LikeCount int  `json:"likeCount"`
	Liked     bool `json:"liked"`
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

	var commentsResult []GetCommentResultItem
	comments, err := content_service.GetComments(form.QuestionId)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	token, err := c.Cookie("access-token")
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	for _, value := range comments {
		likeCount, err := content_service.GetUserLikeCount(map[string]interface{}{
			"typeId":   value.ID,
			"likeType": models.LikeComment,
		})
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}

		liked, err := content_service.GetUserLiked(map[string]interface{}{
			"typeId":    value.ID,
			"likeType":  models.LikeComment,
			"accountId": claims.UserId,
		})

		commentsResult = append(commentsResult, GetCommentResultItem{
			value,
			likeCount,
			liked,
		})
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"comments": commentsResult,
	})
}

type CreateCommentForm struct {
	QuestionId int    `form:"questionId" valid:"Required"`
	Content    string `form:"content" valid:"Required"`
	ParentId   int    `form:"parentId"`
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

	token, err := c.Cookie("access-token")
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	err = content_service.CreateComment(map[string]interface{}{
		"questionId": form.QuestionId,
		"content":    form.Content,
		"parentId":   form.ParentId,
		"authorId":   claims.UserId,
	})

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
