package model

/**
 * 商铺 - 服务 关系表
 */
type ShopService struct {
	ShopId    int64 `gorm:"primaryKey;column:shop_id" json:"shop_id"`
	ServiceId int64 `gorm:"primaryKey;column:service_id" json:"service_id"`
}

func (ShopService) TableName() string {
	return "tb_shop_service"
}
