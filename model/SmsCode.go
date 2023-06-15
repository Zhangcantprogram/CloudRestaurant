package model

type SmsCode struct {
	Id         int64  `gorm:"primaryKey;column:id" json:"id"`
	Phone      string `gorm:"column:phone" json:"phone"`
	BizId      string `gorm:"column:biz_id" json:"biz_id"`
	Code       string `gorm:"column:code" json:"code"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
}

// 自定义表名,与数据库表名一致
func (SmsCode) TableName() string {
	return "tb_smsCode"
}
