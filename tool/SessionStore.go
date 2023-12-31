package tool

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

var Sess sessions.Session

// 初始化session操作，并使用中间件
func InitSession(engine *gin.Engine) {
	redisConfig := GetConfig().Redis
	store, err := redis.NewStore(10, "tcp", redisConfig.Addr+":"+redisConfig.Port, redisConfig.Password, []byte("secret"))
	log.Println("session store ----------->", store)
	if err != nil {
		log.Println(err.Error())
	}

	engine.Use(sessions.Sessions("mysession", store))

}

// 初始化全局session
func InitSess(ctx *gin.Context) {
	Sess = sessions.Default(ctx)
}

// 自定义set方法
func SetSess(session sessions.Session, key interface{}, value interface{}) error {
	//session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

// 自定义get方法
func GetSess(session sessions.Session, key interface{}) interface{} {
	//session := sessions.Default(context)
	if session == nil {
		return nil
	}
	sess := session.Get(key)
	//if sess == nil {
	//	log.Println("key值对应的value为nil")
	//	return nil
	//}
	log.Println("key ---->", key)
	log.Println("session get成功！sess--->", sess)

	return sess

}
