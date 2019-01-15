package validate

import (
	"github.com/gin-gonic/gin"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type ParamUnFoundError struct {
	paramName string
}

func (pue *ParamUnFoundError) Error() string {
	if pue.paramName == "" {
		return "没有找到参数"
	}else{
		return fmt.Sprintf("没有找到参数名%s的参数具体内容",pue.paramName)
	}
}

func newParamUnFoundError(paramName string ) error {
	return &ParamUnFoundError{paramName:paramName}
}

// 对str类型的参数进行验证
func StrNullValidate(paramName string,c *gin.Context) (string,error) {
	param,exist := GetParam(paramName,c)
	if !exist {
		return "",newParamUnFoundError(paramName)
	}else{
		return param,nil
	}
}
// 对int类型的参数进行验证
func IntNullValidate(paramName string,c *gin.Context) (int,error)  {
	param,exist := GetParam(paramName,c)

	if !exist {
		return 0,newParamUnFoundError(paramName)
	}
	paramInt,err := strconv.Atoi(param)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("参数%s,请传入int类型的值",paramName))
	}

	return paramInt,nil
}

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
	DateTimeShortFormat = "2006-1-2 15:4:5"
	DateShortFormat = "2006-1-2"
)

// 对时间类型的参数进行验证
func TimeNullValidate(paramName string,c *gin.Context) (*time.Time,error)  {
	param,exist := GetParam(paramName,c)

	if !exist || param == "" {
		return nil,newParamUnFoundError(paramName)
	}
	var timeInstance time.Time
	var err error
	switch len(param) {
	case 19:
		timeInstance,err = time.Parse(DateTimeFormat,param)
		break
	case 10:
		timeInstance,err = time.Parse(DateFormat,param)
		break
	case 15:
		timeInstance,err = time.Parse(DateTimeShortFormat,param)
		break
	case 8:
		timeInstance,err = time.Parse(DateShortFormat,param)
		break
	default:
		return nil,errors.New("未知的时间格式")
	}
	return &timeInstance,err
}

func GetParam(paramName string,c *gin.Context) (string,bool) {
	param,exist := c.GetPostForm(paramName)
	if !exist {
		param,exist = c.GetQuery(paramName)
	}

	if !exist {
		return "",false
	}else{
		return param,false
	}
}
