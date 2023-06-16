package controller

import (
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)

	//手机号 + 短信验证码登录
	engine.POST("/api/login_sms", mc.smsLogin)

	//获取验证码
	engine.GET("/api/captcha", mc.captcha)

	//验证验证码是否正确
	engine.POST("/api/vertifycha", mc.verifyCaptcha)

	//进行用户名+密码 + 图片验证码 登录功能
	engine.POST("/api/login_pwd", mc.nameLogin)

	// 用户头像上传
	engine.POST("/api/upload/avator", mc.uploadAvatar)

	//用户信息查询
	//engine.GET("/api/userinfo", mc.userInfo)
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
	//初始化全局session
	tool.InitSess(context)

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
		//将member进行序列化，为接下来的
		sess, _ := json.Marshal(member)
		//将登录成功的用户信息保存到session中
		log.Println("member.Id-------------->", member.Id)
		log.Println("sess------------->", sess)
		err = tool.SetSess(tool.Sess, "user_"+strconv.Itoa(int(member.Id)), sess)
		//设置cookie
		//context.SetCookie("cookie_user", strconv.Itoa(int(member.Id)), 10*60, "/", "localhost", true, true)
		if err != nil {
			//说明存在error，即用户虽然登录成功，但是session保存失败，也视为登录失败
			tool.Failed(context, "用户登录失败！")
			log.Println("Setsess用户登录失败!")
			return
		}
		tool.Success(context, member)
		log.Println("手机号+短信验证码 登录成功！！")
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
func (mc *MemberController) verifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, "参数解析失败！......")
		return
	}
	result := tool.VerifyCaptcha(captcha.Id, captcha.VerifyValue)
	fmt.Println("result++----------->", result)
	if result {
		fmt.Println("验证码验证通过！！！")
	} else {
		fmt.Println("验证码验证失败！......")
	}
}

// 根据用户名和密码 来进行登录
func (mc *MemberController) nameLogin(context *gin.Context) {
	tool.InitSess(context)

	//1、获取前端传来的参数
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败！")
		//fmt.Println("参数解析失败!.....")
		return
	}
	//2、验证验证码
	fmt.Println(loginParam)
	result := tool.VerifyCaptcha(loginParam.Id, loginParam.Value)
	//fmt.Println("result----------->", result)
	if !result {
		tool.Failed(context, "验证码错误，请重新输入！")
		return
	}
	//3、完成登录
	ms := service.MemberService{}
	member := ms.NameLogin(loginParam.Name, loginParam.Password)
	if member != nil {
		//将member进行序列化，为接下来保存session做准备
		sess, _ := json.Marshal(member)
		//将登录成功的用户信息保存到session中
		fmt.Println("member.Id-------------->", member.Id)
		//log.Println("sess------------->", sess)
		log.Println("set session user_" + strconv.Itoa(int(member.Id)))
		err = tool.SetSess(tool.Sess, "user_"+strconv.Itoa(int(member.Id)), sess)
		//设置cookie
		//context.SetCookie("cookie_user", strconv.Itoa(int(member.Id)), 10*60, "/", "localhost", true, true)
		if err != nil {
			//说明存在error，即用户虽然登录成功，但是session保存失败，也视为登录失败
			tool.Failed(context, "用户登录失败！")
			return
		}
		tool.Success(context, member)
		log.Println("用户名+密码+验证码 登录成功")
		return
	}
	//否则登录失败
	tool.Failed(context, "用户登录失败！")
}

// 用户头像上传
func (mc *MemberController) uploadAvatar(context *gin.Context) {
	//1、解析上传的参数：user_id，file
	userId := context.PostForm("user_id")
	log.Println("用户登录时的userId---->" + userId)
	file, err := context.FormFile("avatar")
	log.Println("用户登录时的avatar---->", file.Size)
	if err != nil || userId == "" {
		tool.Failed(context, "参数解析失败！")
	}

	//2、判断用户是否已经登录，使用session
	log.Println("get session user_" + userId)
	sess := tool.GetSess(tool.Sess, "user_"+userId)

	if sess == nil {
		tool.Failed(context, "用户登录信息有误！")
		log.Println("用户登录信息有误！")
		return
	}

	//将用户参数进行反序列化
	var member model.Member
	json.Unmarshal(sess.([]byte), &member)
	//3、将file保存到本地
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(context, "头像更新失败！")
		return
	}

	//将文件上传到FastDFS文件分布式系统
	fileId := tool.UploadFile(fileName)
	if fileId != "" {
		//删除本地文件
		//os.Remove(fileName)

		//将文件对应路径保存到数据库中
		ms := service.MemberService{}
		//这里使用 fileName[1:]的切片形式，是为了保存的过程中不保存一开始的那个 “.”
		path := ms.UploadAvatar(member.Id, fileId)
		if path != "" {
			tool.Success(context, tool.FileServerAddr()+"/"+path)
			return
		}
	}

	//5、返回结果
	tool.Failed(context, "头像更新失败！")
}

/**
 * 用户信息查询
 */
func (mc *MemberController) userInfo(context *gin.Context) {
	cookie, err := tool.CookieAuth(context)
	log.Println("cookie---->", cookie)
	if err != nil {
		log.Println(err.Error())
		context.Abort()
		log.Println("cookie获取失败！")
		tool.Failed(context, "还未登录，请先登录！")
	}

	memberService := service.MemberService{}
	member := memberService.GetMemberInfoById(cookie.Value)
	log.Println("cookie---->" + cookie.Value)
	if member != nil {
		tool.Success(context, map[string]interface{}{
			"id":            member.Id,
			"user_name":     member.UserName,
			"mobile":        member.Mobile,
			"register_time": member.RegisterTime,
			"avatar":        member.Avatar,
			"balance":       member.Balance,
			"city":          member.City,
		})
		return
	}
	tool.Failed(context, "用户信息获取失败！")
}
