package model

// 食品类别结构体
type FoodCategory struct {
	//类别ID
	Id int `gorm:"primary_key;column:id" json:"id"`
	//食品类别标题
	Title string `gorm:"column:tile" json:"title"`
	//食品描述
	Description string `gorm:"column:description" json:"description"`
	//食品种类图片
	ImageUrl string `gorm:"column:image_url" json:"image_url"`
	//食品类别连接
	LinkUrl string `gorm:"column:link_url" json:"link_url"`
	//该类别是否在服务状态
	IsInServing int `gorm:"column:is_in_serving" json:"is_in_serving"`
}

// 自定义表名,与数据库表名一致
func (FoodCategory) TableName() string {
	return "food_category"
}
