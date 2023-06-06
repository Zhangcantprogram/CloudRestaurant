package model

// 会员数据 定义结构体
type Member struct {
	Id           int64   `gorm:"primary_key;column:id" json:"id"`
	UserName     string  `gorm:"column:user_name" json:"user_name"`
	Mobile       string  `gorm:"column:mobile" json:"mobile"`
	Password     string  `gorm:"column:password" json:"password"`
	RegisterTime int64   `gorm:"column:register_time" json:"register_time"`
	Avatar       string  `gorm:"column:avatar" json:"avatar"`
	Balance      float64 `gorm:"column:balance" json:"balance"`
	IsActive     int8    `gorm:"column:is_active" json:"is_active"`
	city         string  `gorm:"column:city" json:"city"`
}

// 自定义表名,与数据库表名一致
func (Member) TableName() string {
	return "tb_member"
}
