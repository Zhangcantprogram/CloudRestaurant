package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
)

type GoodsService struct {
}

func (gs *GoodsService) GetGoods(shopId int) []model.Goods {
	//调用dao方法
	goodsDao := dao.GoodsDao{}
	return goodsDao.GetGoods(shopId)
}
