package main

import (
	"crypto/md5"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"system_tpl/business/auth"
	"system_tpl/core"
)

func main() {
	core.InitDB()
	defer core.DB.Close()

	// 初始化用户和权限表
	core.DB.AutoMigrate(
		&auth.User{},
		&auth.Permission{},
	)
	permission := auth.Permission{Name: "管理员"}
	core.DB.Save(&permission)
	passwd := "admin123"
	has := md5.Sum([]byte(passwd))
	passwdMD5 := fmt.Sprintf("%x", has)
	user := auth.User{Name: "admin", PasswordMD5: passwdMD5, Permission: permission}
	core.DB.Save(&user)

}
