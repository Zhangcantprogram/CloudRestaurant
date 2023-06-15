package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

type ShopController struct {
}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.getShops)
	engine.GET("/api/search_shops", sc.searchShop)
}

/**
 * 获取商铺列表
 */
func (sc *ShopController) getShops(context *gin.Context) {
	//1、先从前端获取数据
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")

	if longitude == "" || longitude == "undefined" || latitude == "" || longitude == "undefined" {
		//设置默认经纬度：北京
		longitude = "116.34"
		latitude = "40.34"
	}

	shopService := service.ShopService{}
	shops := shopService.GetShops(longitude, latitude)
	if len(shops) != 0 {
		tool.Success(context, shops)
		return
	}
	tool.Failed(context, "商铺数据返回失败！")
}

// 根据关键字进行搜索
func (sc *ShopController) searchShop(context *gin.Context) {
	//1、先从前端获取数据
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	keyword := context.Query("keyword")

	if keyword == "" {
		tool.Failed(context, "请输入商铺名称！")
		return
	}

	if longitude == "" || longitude == "undefined" || latitude == "" || longitude == "undefined" {
		//设置默认经纬度：北京
		longitude = "116.34"
		latitude = "40.34"
	}

	shopService := service.ShopService{}
	shops := shopService.SearchShops(longitude, latitude, keyword)
	tool.Success(context, shops)
	//if len(shops) != 0 {
	//	tool.Success(context, shops)
	//	return
	//}
	//tool.Failed(context, "商铺数据返回失败！")
}
