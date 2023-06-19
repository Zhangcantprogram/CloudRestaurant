package tool

import (
	"github.com/smartwalle/alipay/v3"
	"log"
)

var (
	client     *alipay.Client
	appId      string
	privateKey string
	publicKey  string
)

func InitAlipay() {
	alipayConfig := GetConfig().Alipay
	appId := alipayConfig.AppId
	privateKey := alipayConfig.PrivateKey
	publicKey := alipayConfig.PublicKey
	client, _ = alipay.New(appId, privateKey, false, alipay.WithSandboxGateway("https://openapi-sandbox.dl.alipaydev.com/gateway.do"))
	//支付宝公钥
	client.LoadAliPayPublicKey(publicKey)
	log.Println("支付宝初始化成功")

}

func WebPageAlipay(amount string) string {
	tradePagePay := alipay.TradePagePay{}
	//订单付款后跳转的网址页面
	//tradePagePay.ReturnURL = "https://www.baidu.com"
	//tradePagePay.NotifyURL = "https://www.baidu.com"

	orderId := "2088722004229570"
	//log.Println("orderId----->", orderId)
	//付款标题
	tradePagePay.Subject = "云餐厅，订单号：" + orderId
	//商家订单号
	tradePagePay.OutTradeNo = orderId
	//价格
	tradePagePay.TotalAmount = amount

	tradePagePay.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(tradePagePay)
	if err != nil {
		log.Println(err)
	}

	//log.Println("url----->", url)

	return url.String()
}

func AliPayNotify(url string) {

}
