package v1

import (
	"api/models"
	"api/pkg/e"
	"api/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Login(c *gin.Context) {
	mobile := c.PostForm("mobile") //取出参数手机号mobile
	vCode := c.PostForm("code")    //取出参数验证码code

	// 对请求参数验证
	validate := validation.Validation{}
	validate.Mobile(mobile, "Mobile").Message("手机号有误")
	validate.Length(vCode, 4, "Code").Message("验证码格式不正确")
	// 校验错误，返回错误信息
	if isOk := checkValidation(&validate, c); isOk == false {
		return
	}

	// 数据库根据手机号查询用户信息
	user, err := models.FindUserByMobile(mobile)
	if gorm.IsRecordNotFoundError(err) {
		util.ResponseWithJson(e.ErrorAuth, "手机号未注册", c)
		return
	} else {
		if err != nil {
			util.ResponseWithJson(e.ERROR, "数据库操作错误", c)
			return
		}
	}

	// 生成token
	token, err := util.GeterateToken(user.ID, user.Mobile)
	if err != nil {
		util.ResponseWithJson(e.ERROR, "创建token失败", c)
		return
	}

	util.ResponseWithJson(e.SUCCESS, gin.H{
		"User":  user,
		"Token": token,
	}, c)
}

func Register(c *gin.Context) {
	mobile := c.PostForm("mobile") //取出参数手机号mobile
	vCode := c.PostForm("code")    //取出参数验证码code

	// 对请求参数验证
	validate := validation.Validation{}
	validate.Mobile(mobile, "Mobile").Message("手机号有误")
	validate.Length(vCode, 4, "Code").Message("验证码格式不正确")
	// 校验错误，返回错误信息
	if isOk := checkValidation(&validate, c); isOk == false {
		return
	}

	// 数据库根据手机号查询用户信息
	user, err := models.FindUserByMobile(mobile)
	if gorm.IsRecordNotFoundError(err) {
		// 查询用户不存在，执行注册逻辑
		user, err = models.CreateUser(mobile)
		if err != nil {
			util.ResponseWithJson(e.ERROR, "用户注册失败", c)
			return
		}
	} else {
		util.ResponseWithJson(e.ERROR, "用户已存在", c)
		return
	}

	// 生成token
	token, err := util.GeterateToken(user.ID, user.Mobile)
	if err != nil {
		util.ResponseWithJson(e.ERROR, "创建token失败", c)
		return
	}

	util.ResponseWithJson(e.SUCCESS, gin.H{
		"User":  user,
		"Token": token,
	}, c)
}

// 参数校验
func checkValidation(valid *validation.Validation, c *gin.Context) bool {
	if valid.HasErrors() {
		var errs string
		for _, err := range valid.Errors {
			errs = err.Message
			break
		}
		util.ResponseWithJson(e.InvalidParams, errs, c)
		return false
	}
	return true
}
