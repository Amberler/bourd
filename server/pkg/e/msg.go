package e

// MsgFlags 编码消息
var MsgFlags = map[int]string{
	SUCCESS:                     "ok",
	ERROR:                       "fail",
	InvalidParams:               "请求参数错误",
	ErrorExistTag:               "已存在该标签名称",
	ErrorNotExistTag:            "该标签不存在",
	ErrorNotExistArticle:        "该文章不存在",
	ErrorAuthCheckTokenFail:     "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:  "Token已超时",
	ErrorAuthToken:              "Token错误",
	ErrorAuth:                   "用户不存在",
	ErrorExistAuth:              "用户已存在",
	ErrorAuthPassword:           "密码错误",
	ErrorUploadSaveImageFail:    "保存图片失败",
	ErrorUploadCheckImageFail:   "检查图片失败",
	ErrorUploadCheckImageFormat: "校验图片错误，图片格式或大小有问题",
}

// GetMsg 根据传入的编码。获取对应的编码消息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
