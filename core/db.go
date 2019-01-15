package core

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gw_system/config"
)

var DB *gorm.DB

func InitDB()  {
	db,err := gorm.Open(config.Database["ENGINE"].(string), fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database["USER"],config.Database["PASSWORD"],config.Database["HOST"],config.Database["PORT"],config.Database["NAME"]))

	if err!=nil {
		panic(err)
	}
	DB = db
}