package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"log"
)

type FoodCategoryService struct {
}

// 获取食品分类的全部数据
func (fcs *FoodCategoryService) GetAll() ([]model.FoodCategory, error) {
	fcd := dao.FoodCategoryDao{}
	categories, err := fcd.GetAll()
	if err != nil {
		//说明出现了错误
		log.Println(err.Error())
		return nil, err
	}
	return categories, nil
}
