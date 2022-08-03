package utils

import "pascal_practice_server/pkg/setting"

func SetUp() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
