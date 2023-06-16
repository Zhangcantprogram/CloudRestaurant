package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type GoodsController struct {
}

func (gc *GoodsController) Router(engine *gin.Engine) {
	engine.GET("/api/goods", gc.getGoods)
	engine.DELETE("/api/delete_goods", gc.deleteGoodsById)
}

// 获取对应商家的在售商品
func (gc *GoodsController) getGoods(context *gin.Context) {
	id, exist := context.GetQuery("shop_id")
	if !exist {
		tool.Failed(context, "参数解析失败！")
		return
	}

	shopId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		tool.Failed(context, "参数解析失败！")
		return
	}

	//调用sevice方法
	goodsService := service.GoodsService{}
	goods := goodsService.GetGoods(shopId)
	if len(goods) != 0 {
		tool.Success(context, goods)
		return
	}

	tool.Failed(context, "未查询到商品信息！")
}

// 根据Id删除goods
func (gc *GoodsController) deleteGoodsById(context *gin.Context) {
	id := context.Query("goods_id")

	goodsService := service.GoodsService{}
	goodsService.DeleteGoodsById(id)

	tool.Success(context, "删除成功！")
}
