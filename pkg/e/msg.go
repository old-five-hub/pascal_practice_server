package e

var MsgFlags = map[int]string{
	SUCCESS:             "ok",
	ERROR:               "fail",
	ERROR_GET_TAGS_FAIL: "获取标签失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
