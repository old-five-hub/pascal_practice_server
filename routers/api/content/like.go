package content

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/e"
	"pascal_practice_server/pkg/utils"
	"pascal_practice_server/service/content_service"
)

type UpdateLikeForm struct {
	TypeId     int `form:"typeId" valid:"Required"`
	LikeType   int `form:"likeType" valid:"Range(0, 1)"`
	LikeStatus int `form:"likeStatus" valid:"Range(0, 1)"`
}

func UpdateLike(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = UpdateLikeForm{}
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

	userLikeData := map[string]interface{}{
		"typeId":     form.TypeId,
		"likeType":   form.LikeType,
		"accountId":  claims.UserId,
		"likeStatus": form.LikeStatus,
	}

	existUserLike, err := content_service.ExistUserLike(userLikeData)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if existUserLike.ID > 0 {
		if existUserLike.LikeStatus == form.LikeStatus {
			appG.Response(http.StatusBadRequest, e.ERROR, nil)
			return
		}

		if err := content_service.UpdateUserLike(existUserLike.ID, userLikeData); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}
	} else {
		if err := content_service.CreateUserLike(userLikeData); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}
	}

	userLikeStatData := map[string]interface{}{
		"typeId":   form.TypeId,
		"likeType": form.LikeType,
	}

	existUserLikeStat, err := content_service.ExistUserLikeStat(userLikeStatData)

	if existUserLikeStat.ID > 0 {
		var diff int
		if form.LikeStatus == 0 {
			diff = -1
		} else {
			diff = 1
		}

		if err := content_service.UpdateUserLikeStat(existUserLikeStat.ID, map[string]interface{}{
			"typeId":    form.TypeId,
			"likeType":  form.LikeType,
			"likeCount": existUserLikeStat.LikeCount + diff,
		}); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}
	} else {
		if err := content_service.CreateUserLikeStat(map[string]interface{}{
			"typeId":    form.TypeId,
			"likeType":  form.LikeType,
			"likeCount": form.LikeStatus,
		}); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
