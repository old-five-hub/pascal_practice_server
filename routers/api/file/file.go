package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"pascal_practice_server/pkg/app"
	"pascal_practice_server/pkg/cos"
	"pascal_practice_server/pkg/e"
	"pascal_practice_server/pkg/setting"
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
		"uri": path,
	})
}

func GetFileContent(path string, dir string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", setting.TencentSetting.CosUrl, dir, path)

	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", body), nil
}
