package dao

import (
	"CloudRestaurant/global"
	"CloudRestaurant/model"
	"log"
)

type MemberDao struct {
}

// 向数据库tb_smsCode表中插入一条数据
func (md *MemberDao) InsertSmsCode(sms model.SmsCode) int64 {
	result := global.DB.Model(&model.SmsCode{}).Create(&sms)
	if result.Error != nil {
		//如果有错误则打印错误
		log.Fatalln(result.Error.Error())
	}
	//返回受影响的行数
	return result.RowsAffected
}

// 验证手机号和验证码是否存在
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var smsCode model.SmsCode
	find := global.DB.Debug().Where("phone = ? and code = ? ", phone, code).First(&smsCode)
	if smsCode.Id == 0 {
		//如果没查询出来数据，则打印错误
		log.Fatalln(find.Error.Error())
	}
	return &smsCode
}

// 根据phone在member表中查询数据
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	find := global.DB.Debug().Where("mobile = ?", phone).First(&member)
	if member.Id == 0 {
		//如果有错误则打印错误
		log.Fatalln(find.Error.Error())
	}
	return &member
}

// 在member表中新添加一个数据
func (md *MemberDao) InsertMember(member model.Member) int64 {
	result := global.DB.Create(&member)
	if result.Error != nil {
		//如果有错误则打印错误
		log.Fatalln(result.Error.Error())
	}
	//返回受影响的行数
	return result.RowsAffected
}
