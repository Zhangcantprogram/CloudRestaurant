package model

type Service struct {
	//id 自增长
	Id int `gorm:"primaryKey;column:id" json:"id"`
	//服务名称
	Name string `gorm:"column:name" json:"name"`
	//服务描述
	Description string `gorm:"column:description" json:"description"`
	//图标名称
	IconName string `gorm:"column:icon_name" json:"icon_name"`
	//图标颜色
	IconColor string `gorm:"column:icon_color" json:"icon_color"`
}

func (Service) TableName() string {
	return "tb_service"
}
