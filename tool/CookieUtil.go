package tool

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const CookieName string = "cookie_user"
const CookieTimeLength = 10 * 60 //10分钟，单位为秒

func CookieAuth(context *gin.Context) (*http.Cookie, error) {
	cookie, err := context.Request.Cookie(CookieName)
	if err == nil {
		//说明没有错误，则继续设置cookie
		log.Println("cookie.Name ----->", cookie.Name)
		context.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	} else {
		log.Println("CookieAuth---->", err.Error())
		return nil, err
	}
	return cookie, nil
}
