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
func (sd *ShopDao) GetShops(longitude, latitude float64, keyword string) []model.Shop {
	var shops []model.Shop
	log.Println("dao keyword -----> ", keyword)
	if keyword == "" {
		//则说明没有使用条件查询
		tool.DB.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and status = 1",
			longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
		//find := tool.DB.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and status = 1",
		//	longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
		//if find.Error != nil && find.RowsAffected!=0{
		//	log.Println(find.Error.Error())
		//	return nil
		//}
	} else {
		//说明需要进行关键字搜索
		log.Println("进行关键字搜索")
		tool.DB.Debug().Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1",
			longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, "%"+keyword+"%").Find(&shops)
		//find := tool.DB.Debug().Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1",
		//	longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, keyword).Find(&shops)
		//if find.Error != nil {
		//	log.Println(find.Error.Error())
		//	return nil
		//}
	}

	return shops

}
