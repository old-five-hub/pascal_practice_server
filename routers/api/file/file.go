package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/cos"
	"pascal_practice_server/pkg/e"
	path2 "path"
)

const (
	Image    string = "image"
	Question        = "question"
)

func UploadFile(c *gin.Context) {
	appG := app.Gin{C: c}
	file, header, err := appG.C.Request.FormFile("file")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_FILE, nil)
		return
	}

	ext := path2.Ext(header.Filename)
	cosType := appG.C.PostForm("cosType")

	fmt.Println(cosType)

	if cosType == "" || (cosType != Image && cosType != Question) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_FILE_UNKNOW_TYPE, nil)
		return
	}

	path, err := cos.UploadFile(cosType, file, ext)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_FILE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		path: path,
	})
}
