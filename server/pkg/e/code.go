package e

const (
	SUCCESS       = 200 //成功响应请求
	ERROR         = 500 //错误响应请求
	InvalidParams = 400 //请求参数无效

	ErrorExistTag        = 410 //标记错误
	ErrorNotExistTag     = 411 //错误的不存在的标记
	ErrorNotExistArticle = 412 //错误的不存在的文章

	ErrorAuthCheckTokenFail    = 413 //token无效
	ErrorAuthCheckTokenTimeout = 414 //token超时
	ErrorAuthToken             = 415 //token错误
	ErrorAuth                  = 416 //无效的用户
	ErrorExistAuth             = 417 //手机号已存在
	ErrorAuthPassword          = 418 //密码错误

	ErrorUploadSaveImageFail    = 460 // 保存图片失败
	ErrorUploadCheckImageFail   = 461 // 检查图片失败
	ErrorUploadCheckImageFormat = 462 // 校验图片错误，图片格式或大小有问题
)
