package controller

import (
	"CloudRestaurant/model"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
	"log"
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

	shopService := &service.ShopService{}
	shops := shopService.SearchShops(longitude, latitude, keyword)
	if len(shops) == 0 {
		tool.Failed(context, "暂未获得商铺信息！")
		return
	}
	for i := range shops {
		services := shopService.GetServiceByShopId(int(shops[i].Id))
		if len(services) > 0 {
			shopSupports := make([]model.Service, len(services))
			copy(shopSupports, services)
			shops[i].Supports = shopSupports
		}
	}
	log.Println("shops", shops)
	tool.Success(context, shops)
}
