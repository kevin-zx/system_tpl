package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gw_system/model"
	"fmt"
	"crypto/md5"
)

func main() {
	movieDb, err := gorm.Open("mysql", "remote:Iknowthat@tcp(115.159.79.85:3306)/gw?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer movieDb.Close()
	movieDb.AutoMigrate(
		&model.User{},
		&model.Permission{},
	)
	permission := model.Permission{Name:"管理员"}
	movieDb.Save(&permission)
	passwd := "admin123"
	has := md5.Sum([]byte(passwd))
	passwdMD5 := fmt.Sprintf("%x",has)
	user := model.User{Name:"admin",PasswordMD5:passwdMD5,Permission:permission}
	movieDb.Save(&user)

}


