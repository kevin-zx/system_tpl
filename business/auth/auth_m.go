package auth

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string     `gorm:"unique" json:"user_name"`
	PasswordMD5  string     `json:"password_md5"`
	Permission   Permission `json:"permission"`
	PermissionID uint
}

type Permission struct {
	gorm.Model
	Name string `json:"name"`
}
