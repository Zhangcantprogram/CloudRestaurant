package global

import (
	"github.com/jinzhu/gorm"
)

// 定义gorm数据库驱动的全局变量
var DB *gorm.DB = nil

// 定义redis驱动的全局变量
//var Redis *redis.Client
