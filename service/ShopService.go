package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"log"
	"strconv"
)

type ShopService struct {
}

// 获取商铺信息
func (ss *ShopService) GetShops(longitude, latitude string) []model.Shop {
	lon, err := strconv.ParseFloat(longitude, 10)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	lat, err := strconv.ParseFloat(latitude, 10)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	shopDao := dao.ShopDao{}
	return shopDao.GetShops(lon, lat)
}
