package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	cfg, err := tool.PareConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()

	//启动gorm连接数据库，并将DB保存为全局变量
	tool.GetDb()
	defer tool.DB.Close()

	//初始化redis配置
	tool.InitRedisStore()

	//初始化session
	tool.InitSession(app)

	//初始化支付宝功能
	tool.InitAlipay()

	//设置全局跨域访问
	app.Use(Cors())

	//注册路由
	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}

// 路由配置
func registerRouter(engine *gin.Engine) {
	new(controller.HelloController).Router(engine)
	new(controller.MemberController).Router(engine)
	new(controller.FoodCategoryController).Router(engine)
	new(controller.ShopController).Router(engine)
	new(controller.GoodsController).Router(engine)
	new(controller.AliPayController).Router(engine)
}

// 跨域中间件配置
// 跨域访问：cross  origin resource share
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}

		//处理请求
		context.Next()
	}
}
