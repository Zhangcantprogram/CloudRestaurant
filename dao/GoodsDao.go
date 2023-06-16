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

func (gd GoodsDao) DeleteGoodsById(goodsId int) {
	var goods model.Goods
	//tool.DB.Where("id = ?", goodsId).First(&goods)
	//tool.DB.Delete(&goods)
	tool.DB.Where("id = ?", goodsId).Delete(&goods)
}
