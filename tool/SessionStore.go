package tool

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

// 初始化session操作，并使用中间件
func InitSession(engine *gin.Engine) {
	redisConfig := GetConfig().Redis
	store, err := redis.NewStore(10, "tcp", redisConfig.Addr+":"+redisConfig.Port,
		redisConfig.Password, []byte("secret"))
	log.Println("session store ----------->", store)
	if err != nil {
		log.Println(err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
}

// 自定义set方法
func SetSess(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)

	if session == nil {
		return nil
	}

	session.Set(key, value)
	//log.Println("session设置成功！", session.Get(key))
	session.Save()
	sessionId := session.ID()
	//sess := session.Get(key)
	//log.Println("session get成功！sess--->", sess)
	log.Println("sessionId---->", sessionId)

	return session.Save()
}

// 自定义get方法
func GetSess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	sess := session.Get(key)
	if sess == nil {
		log.Println("key值对应的value为nil")
		return nil
	}
	log.Println("key ---->", key)
	log.Println("session get成功！sess--->", sess)

	return session.Get(key)

}
