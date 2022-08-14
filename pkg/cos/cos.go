package cos

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"pascal_practice_server/pkg/setting"
)

var Client *cos.Client

func SetUp() {
	u, _ := url.Parse(setting.TencentSetting.CosUrl)
	b := &cos.BaseURL{BucketURL: u}
	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  setting.TencentSetting.SecretId,
			SecretKey: setting.TencentSetting.SecretKey,
		},
	})
}

func UploadFile(path string, file multipart.File, ext string) (string, error) {
	u := uuid.NewV4()

	name := fmt.Sprintf(`%s/%s%s`, path, u, ext)

	_, err := Client.Object.Put(context.Background(), name, file, nil)
	if err != nil {
		return "", err
	}

	return name, nil
}
