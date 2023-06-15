package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type MemberService struct {
}

// 完成 手机号+验证码  登录服务
func (ms *MemberService) SmsLogin(loginParam param.SmsLoginParam) *model.Member {
	//todo
	//1.获取到手机号和验证码

	//2.验证手机号+验证码是否正确
	md := dao.MemberDao{}
	smsCode := md.ValidateSmsCode(loginParam.Phone, loginParam.Code)
	//fmt.Println(smsCode.Phone)
	if smsCode.Id == 0 {
		return nil
	}
	//3.根据手机号在member表中查询记录
	member := md.QueryByPhone(loginParam.Phone)
	//fmt.Println(member.Mobile)
	if member != nil {
		//说明能在member表中查询到数据
		return member
	}
	//4.新创建一个member记录，并保存
	newMember := model.Member{UserName: loginParam.Phone, Mobile: loginParam.Phone, RegisterTime: time.Now().Unix()}
	row := md.InsertMember(newMember)
	if row != 0 {
		return &newMember
	}
	return nil
}

// 发送短信服务
func (ms *MemberService) SendCode(phone string) bool {
	//1.产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//2.调用阿里云sdk完成发送
	smsConfig := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(smsConfig.RegionId, smsConfig.AppKey, smsConfig.AppSecret)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	//设置请求协议
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = smsConfig.SignName
	request.TemplateCode = smsConfig.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	//3.接收返回结果，并判断发送状态
	//短信验证码发送成功
	if response.Code == "OK" {
		//将验证码保存到数据库中
		smsCode := model.SmsCode{Phone: phone, Code: code, BizId: response.BizId, CreateTime: time.Now().Unix()}
		memberDao := dao.MemberDao{}
		row := memberDao.InsertSmsCode(smsCode)
		return row > 0
	}

	return false
}

// 用户名+密码 实现登录
func (ms *MemberService) NameLogin(name string, password string) *model.Member {
	//1、使用用户名+密码 查询用户信息  如果用户存在，则直接返回member结构体
	md := dao.MemberDao{}
	member := md.QueryByName(name, password)
	if member.Id != 0 {
		//说明该用户已存在
		return member
	}

	//2、如果用户不存在，则作为新用户保存到数据库中
	newMember := model.Member{}
	newMember.UserName = name
	//利用sha256进行加密处理
	newMember.Password = tool.EncoderSha256(password)
	newMember.RegisterTime = time.Now().Unix()
	//再调用InsertMember方法进行用户的添加
	result := md.InsertMember(newMember)
	newMember.Id = result
	return &newMember
}

// 保存文件的本地路径，保存到member表中的avatar字段中
func (ms *MemberService) UploadAvatar(userId int64, fileName string) string {
	md := dao.MemberDao{}
	result := md.UploadAvatarByMember(userId, fileName)
	if result == 0 {
		log.Println("数据库头像字段更新失败！")
		return ""
	}
	return fileName
}

// 获取用户信息
func (ms *MemberService) GetMemberInfoById(userId string) *model.Member {
	id, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	memberDao := dao.MemberDao{}
	return memberDao.GetMemberInfoById(id)
}
