package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"log"
)

type ShopDao struct {
}

const DEFAULT_RANGE = 5

/**
 * 获取商铺信息
 */
func (sd *ShopDao) GetShops(longitude, latitude float64) []model.Shop {
	var shops []model.Shop
	find := tool.DB.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?",
		longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
	if find.Error != nil {
		log.Println(find.Error.Error())
		return nil
	}
	return shops

}
