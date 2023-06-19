package controller

import (
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AliPayController struct {
}

func (apc *AliPayController) Router(engine *gin.Engine) {
	engine.GET("/api/pay", apc.pay)
	//engine.GET("/api/return", apc.notify)
}

// 支付功能
func (apc *AliPayController) pay(context *gin.Context) {
	amount := context.Query("amount")
	//调用支付宝接口
	payUrl := tool.WebPageAlipay(amount)
	if payUrl == "" {
		tool.Failed(context, "支付产生错误！")
		return
	}
	context.Redirect(http.StatusFound, payUrl)
}

// 接受通知接口
func (apc *AliPayController) notify(context *gin.Context) {
	tool.Success(context, "支付成功！")
}
