package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"log"
	"strconv"
)

type GoodsService struct {
}

func (gs *GoodsService) GetGoods(shopId int) []model.Goods {
	//调用dao方法
	goodsDao := dao.GoodsDao{}
	return goodsDao.GetGoods(shopId)
}

// 根据Id删除goods
func (gs *GoodsService) DeleteGoodsById(id string) {
	goodsId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	goodsDao := dao.GoodsDao{}
	goodsDao.DeleteGoodsById(goodsId)
}
