package controller

import (
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)

	engine.POST("/api/login_sms", mc.smsLogin)

	//获取验证码
	engine.GET("/api/captcha", mc.captcha)

	//验证验证码是否正确
	engine.POST("/api/vertifycha", mc.vertifyCaptcha)
}

// http://localhost:8090/api/sendcode?phone=13112345678
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Failed(context, "手机号数据解析失败！")
		return
	}
	ms := service.MemberService{}
	isSend := ms.SendCode(phone)
	if isSend {
		tool.Success(context, "手机验证码发送成功！")
		return
	}

	tool.Failed(context, "手机验证码发送失败！......")
}

// 手机号+短信验证码  登录
func (mc *MemberController) smsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam
	//将请求体里的json通过解码解析出来
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败！......")
	}

	//完成手机号+验证码  登录服务
	sc := service.MemberService{}
	member := sc.SmsLogin(smsLoginParam)
	if member != nil {
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "手机号验证码登录失败！......")
}

// 生成验证码
func (mc *MemberController) captcha(context *gin.Context) {
	//todo 生成验证码并返回客户端
	tool.GenerateCaptcha(context)
}

// 验证验证码是否正确
func (mc *MemberController) vertifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, "参数解析失败！......")
		return
	}
	result := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证码验证通过！！！")
	} else {
		fmt.Println("验证码验证失败！......")
	}
}
