package services

import (
	"crypto/md5"
	"fmt"
	"gw_system/model"
	"encoding/json"
	"gw_system/core"
	"gw_system/config"
)

func MD5(str string) string {
	has :=md5.Sum([]byte(str))
	return fmt.Sprintf("%x", has)
}

func CheckUser(data []byte) map[string]interface{} {
	user := model.User{}
	err := json.Unmarshal(data, &user)

	if err != nil {
		return map[string]interface{}{"code":config.SystemErrCode,"message":"出现错误，请联系管理员"}
	}
	dbUser := model.User{}
	err = core.DB.Where("name = ?",user.Name).First(&dbUser).Error
	if err != nil {

		return map[string]interface{}{"code":config.SystemErrCode,"message":"出现错误，请联系管理员"}
	}
	if dbUser.PasswordMD5 != user.PasswordMD5 {
		return map[string]interface{}{"code":config.SystemErrCode,"message":"用户名和密码不匹配"}
	}

	return map[string]interface{}{"code":config.SuccessCode,"message":"登录成功", "data":map[string]string{"token":MD5(dbUser.PasswordMD5)}}
}