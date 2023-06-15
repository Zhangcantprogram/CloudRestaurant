package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type GoodsDao struct {
}

func (gd *GoodsDao) GetGoods(shopId int) []model.Goods {
	var goods []model.Goods
	tool.DB.Where("shop_id = ?", shopId).Find(&goods)
	return goods
}
