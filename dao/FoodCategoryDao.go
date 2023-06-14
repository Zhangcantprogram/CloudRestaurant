package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"log"
)

type FoodCategoryDao struct {
}

//func NewFoodCategoryDao() *FoodCategoryDao {
//	return &FoodCategoryDao{}
//}

// 查询食品分类的全部数据
func (fcd *FoodCategoryDao) GetAll() ([]model.FoodCategory, error) {
	var foodCategories []model.FoodCategory
	find := tool.DB.Model(&model.FoodCategory{}).Find(&foodCategories)
	if find.Error != nil {
		//如果出现了错误
		log.Println(find.Error.Error())
		return nil, find.Error
	}
	return foodCategories, nil
}
