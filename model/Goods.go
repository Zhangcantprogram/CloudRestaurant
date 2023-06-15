package model

// 食品结构体定义
type Goods struct {
	Id          int64   `gorm:"primaryKey;column:id;AUTO_INCREMENT" json:"id"` //商品Id
	Name        string  `gorm:"column:name" json:"name"`                       //商品名称
	Description string  `gorm:"column:description" json:"description"`         //商品描述
	Icon        string  `gorm:"column:icon" json:"icon"`                       //商品图标
	SellCount   int64   `gorm:"column:sell_count" json:"sell_count"`           //销售份数
	Price       float32 `gorm:"column:price" json:"price"`                     //销售价格
	OldPrice    float32 `gorm:"column:old_price" json:"old_price"`             //原价
	ShopId      int64   `gorm:"column:shop_id" json:"shop_id"`                 //商品ID，表明该商品属于哪个商家
}

func (Goods) TableName() string {
	return "tb_goods"
}
