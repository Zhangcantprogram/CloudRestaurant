package model

/**
 * 商铺信息
 */
type Shop struct {
	//id
	Id int64 `gorm:"primaryKey" json:"id"`
	//商铺名称
	Name string `gorm:"column:name" json:"name"`
	//宣传信息
	PromotionInfo string `gorm:"column:promotion_info" json:"promotion_info"`
	//地址
	Address string `gorm:"column:address" json:"address"`
	//联系电话
	Phone string `gorm:"column:phone" json:"phone"`
	//店铺营业状态
	Status int `gorm:"column:status" json:"status"`

	//经度
	Longitude float64 `gorm:"" json:"longitude"`
	//纬度
	Latitude float64 `gorm:"" json:"latitude"`
	//店铺图标
	ImagePath string `gorm:"column:image_path" json:"image_path"`

	//
	IsNew bool `gorm:"column:is_new" json:"is_new"`
	//
	IsPremium bool `gorm:"column:is_premium" json:"is_premium"`

	//商铺评分
	Rating float32 `gorm:"column:rating" json:"rating"`
	//评分总数
	RatingCount int64 `gorm:"column:rating_count" json:"rating_count"`
	//当前订单总数
	RecentOrderNum int64 `gorm:"column:recent_order_num" json:"recent_order_num"`

	//配送起送价
	MinimumOrderAmount int32 `gorm:"column:minimum_order_amount" json:"minimum_order_amount"`
	//配送费
	DeliveryFee int32 `gorm:"column:delivery_fee" json:"delivery_fee"`

	//营业时间
	OpeningHours string `gorm:"column:opening_hours" json:"opening_hours"`

	//与tb_service表进行链表查询
	//`gorm:"-"`  表示不映射这个字段
	Supports []Service `gorm:"-" json:"supports"`
}

func (Shop) TableName() string {
	return "tb_shop"
}
