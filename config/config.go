package config

//数据库配置
var Database  = map[string]interface{}{
	"ENGINE":"mysql",
	
	//"HOST":"localhost",
	"PORT":3306,
	"NAME":"system", //数据库名称
	"USER":"remote",
	//"USER":"root",
	"PASSWORD":"123456",
}
const SuccessCode = 20000
const SystemErrCode = 20005
const CheckErrCode = 20001
const Port  = 8088
//const Port  = 8080
//// 模板文件目录
//var TemplateDir = "/template/"
//
//// 静态文件目录
//var StaticDir = "/static/"

