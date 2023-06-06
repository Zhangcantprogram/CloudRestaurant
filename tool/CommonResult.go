package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义请求状态码常量
const (
	SUCCESS int = 0 //操作成功
	FAIL    int = 1 //操作失败
)

// 请求成功
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "请求成功！！！",
		"data": v,
	})
}

// 请求失败
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAIL,
		"msg":  "请求失败！......",
		"data": v,
	})
}
