package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
	"log"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	//美食类别
	engine.GET("/api/food_category", fcc.foodCategory)
}

// 获取全部食品分类数据
func (fcc *FoodCategoryController) foodCategory(context *gin.Context) {
	fcs := service.FoodCategoryService{}
	categories, err := fcs.GetAll()
	if err != nil {
		log.Println(err.Error())
		tool.Failed(context, "食品分类数据查询失败！")
	}
	tool.Success(context, categories)
}
