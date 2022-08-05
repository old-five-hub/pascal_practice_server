package account

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/e"
	"pascal_practice_server/pkg/utils"
	"pascal_practice_server/service/account_service"
)

type account struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	a := account{}
	c.BindJSON(&a)
	ok, _ := valid.Valid(&a)

	if !ok {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	accountService := account_service.Account{Username: a.Username, Password: utils.EncodeMD5(a.Password)}

	account, err := accountService.Login()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if account.ID == 0 {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := utils.GenerateToken(a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"token":    token,
		"avatar":   account.Avatar,
		"nickname": account.Nickname,
		"follow":   account.Follow,
	})
}
