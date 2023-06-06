package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.hello)
}

func (hello *HelloController) hello(context *gin.Context) {

}
