package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "账户名或者密码错误请重试",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:           "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:            "标签id不存在",
	ERROR_GET_TAGS_FAIL:            "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:           "统计标签失败",
	ERROR_ADD_TAG_FAIL:             "新增标签失败",
	ERROR_EDIT_TAG_FAIL:            "修改标签失败",
	ERROR_DELETE_TAG_FAIL:          "删除标签失败",
	ERROR_GET_QUESTION_FAIL:        "获取问题列表失败",
	ERROR_GET_QUESION_DETAIL_FAIL:  "获取问题详情失败",

	ERROR_UPLOAD_FILE:             "上传文件失败",
	ERROR_UPLOAD_FILE_UNKNOW_TYPE: "请传入正确的文件类型",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
