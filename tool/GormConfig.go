package tool

import (
	"CloudRestaurant/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDb() {
	var err error
	config := GetConfig().DataBase
	//"root:zsj20020818@tcp(43.139.78.177:33061)/cloudRestaurant?charset=utf8&parseTime=True&loc=Local"
	conn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" +
		config.DbName + "?charset=" + config.Charset + "&parseTime=" + config.ParseTime + "&loc=" + config.Loc
	global.DB, err = gorm.Open(config.Driver, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(global.DB)
	fmt.Println("cloudRestaurant数据库连接成功！！！")
	//return db
	//defer db.Close()
}
