package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"crypto/md5"
	"system_tpl/model"
	"system_tpl/core"
)

func main() {
	core.InitDB()
	defer core.DB.Close()

	// 初始化用户和权限表
	core.DB.AutoMigrate(
		&model.User{},
		&model.Permission{},
	)
	permission := model.Permission{Name:"管理员"}
	core.DB.Save(&permission)
	passwd := "admin123"
	has := md5.Sum([]byte(passwd))
	passwdMD5 := fmt.Sprintf("%x",has)
	user := model.User{Name:"admin",PasswordMD5:passwdMD5,Permission:permission}
	core.DB.Save(&user)

}


