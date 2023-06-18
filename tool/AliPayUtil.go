package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"log"
)

var (
	appID      = "9021000122693054" //支付宝ID
	privateKey = "MIIEpQIBAAKCAQEA6syOH5og3NPNr8DMb+4wACF/Vm+QeHzAOPMSE/RZ6riHwKEivo9Ca7W4mLVd2CSrqFAGUZXYfTTdt6c7KZEe7ZAGBnSLigVl0psWR3/rRXhX+bbhe1ftz4IkRCWsj3T2MdmVsi+Y9WCWFj5YT2B94Rqg0oJ5DC8ZBfad9dYZijFRuD0HKhcteLOL2DUfdRCZALo4hC40qEGFPit2nraqjfNJOEyS59RJtwoCP2Y1wygcrl2Syd9Yp4AxSat/HBKDOB0mo8B8y+PGouk1q0lkntegw9ofC50vtMjnYWnPit5yJ2l/jS6iBYVh88+sPwzjGLGuoJXos0caVJ48dBp7YwIDAQABAoIBABEMz5k94PsIKlyD3JSYMEYiKJQHP+9v56l7BSoLyWfBBKKZUOOTlZgqtv/n06u5Zfmv7fWIsxLMfvkyHJq1HhDUL4vSdLuODsnPVzvT7yq5RXIttIv3Y5547ddBARndvW77ROKOkEXpQJzW5CTHdfydqWX/8XhQBzomoWvTo0EklhFc6FOYuq0oklHmqUJVzpUTG2QYY0K9mqN5UXLCZNYYSMfOz5f4/pzssILX7jPykcWLWVt6mfJuamBDEKbCi3Oe0OssQVq3q5xE9rKANR4PJe0fTnEzigDPZT6CDN/Gi2rgvub0E+eg1LzHcPMsf9F/a8GyMlxt/GG8OsnBFjECgYEA/UvRjTTk/SF+HZphkRF9Oi+ZQSesD4+zFw+nD+PLWeLVfSQPRs4mQ2r5DstSEuq5QyntWb5K9DLGWNWxZ3p/JlUhRVRDhaAasrbpHYtgwoNmGhbHnYaOF0CeaAGEcamrawlLDZuD1f5x6iujqo72Uk/nowSd9rLYtmyOty+QyYsCgYEA7U4wiUL6OshSWG4UXjDYcd2t4N12Lwpqh7qmWUccnG+rOaZqn32mI6wfuHfEQFSZXHFjo/6eKOfStUo1wNGn5nrTnTWavQXAqyL/4X/84T7g7okSh4RqGhpBu654r0cPddq1LzQMcEXd0RlV7P+R2+7clxZ1EHfwr8T0b0Rv4IkCgYEAgmwUtUGkGtaxCmsgi8LWmSxDHDJiPQz/6QLtQZhIb7sFtE7p8spZF2OwZDa6xvTHedbP2OLL01uFl95IX4DKkaFJ5VnS/q4GsTjPZaALrXxdCVfrZqgO59gk+Ga3nJMBqbZ6R8Joqro+S5Y2yq6cFJwCKIVLFOjX/258OcSrJ6sCgYEAgpqG4fSHzXRUxbCH6zMM3NY5jUm1crxT1zGlvfsCRyK6ZPcvNWos1vMzXg4kZ5dn9FW/lhFbMH4uHHKkVOPOeW3eK2bGLAQPfuC7XXtYTQuTNhUgqo746jbLOKUUbLs6Sg29rCYENoJtS1ibTt11hErZAxpsAU1RnRtTwC01OakCgYEAxDGJ3Msmv6IfKdnvZSarlsPsVyaHCMY12BTRPxSw22qceruqPkbous/R5lMCN6JcPLlMYYJ4bkZeX4IYSeyBJhiW28gyrNdUiy4cLahF4kO+ZREIG/10Es2czG89L4grlJTQO2Cjld7yr0ocMEGNs4nDqfcdkgP/3FFn1gibYJ0="
	publicKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuTajJ/+F6+s8S2ZL37VC0WVpitOV5xaTdgg9l6iW/xFjPZbHOciYGo15ySngQZlTjwK8EWSys71lkWglKAEExLG8YfnaZdIRtlCwM+dTEGxsrhd1skGoOnslag8Se7g2AE2GYHSpmrNA145HPZEC33garx+r3yY9i3SOMws28anlHoGFz1XZwORp62L4d1mVKH2Okye6ir6z107rNRgpcIP2ev2yh4GZ/YQm4u+2Wte3HFJTvJVS1f6RzBWQRn+HAowNc+AIcQ6WLG4Zz8ULco8ac4PiX2RzjcjWo9vlVgFQnrKmXfZNkp52f7hxGHQRU0pPMkvUxiVdC1f1gcXFAQIDAQAB"
	client, _  = alipay.New(appID, privateKey, false, alipay.WithSandboxGateway("https://openapi-sandbox.dl.alipaydev.com/gateway.do"))
)

func InitAlipay() {
	//支付宝公钥
	client.LoadAliPayPublicKey(publicKey)
	log.Println("支付宝初始化成功")

}

func WebPageAlipay(ctx *gin.Context) string {
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
	tradePagePay.TotalAmount = "100"

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
