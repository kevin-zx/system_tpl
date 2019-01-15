package auth

import (
	"crypto/md5"
	"fmt"
	"encoding/json"
	"system_tpl/config"
	"system_tpl/core"
)

func MD5(str string) string {
	has :=md5.Sum([]byte(str))
	return fmt.Sprintf("%x", has)
}

// 对用户进行验证
func CheckUser(data []byte) map[string]interface{} {
	user := User{}
	err := json.Unmarshal(data, &user)

	if err != nil {
		return map[string]interface{}{"code":config.SystemErrCode,"message":"出现错误，请联系管理员"}
	}
	dbUser := User{}
	err = core.DB.Where("name = ?",user.Name).First(&dbUser).Error
	if err != nil {

		return map[string]interface{}{"code":config.SystemErrCode,"message":"出现错误，请联系管理员"}
	}
	if dbUser.PasswordMD5 != user.PasswordMD5 {
		return map[string]interface{}{"code":config.SystemErrCode,"message":"用户名和密码不匹配"}
	}

	return map[string]interface{}{"code":config.SuccessCode,"message":"登录成功", "data":map[string]string{"token":MD5(dbUser.PasswordMD5)}}
}