package core

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type BaseController struct {

}
func (bc *BaseController) GetJsonData(c *gin.Context) (*gjson.Result,error) {
	rd,err :=c.GetRawData()
	if err != nil {

		return nil,err
	}
	taskJson := gjson.Parse(string(rd))
	return &taskJson,nil
}

func (bc *BaseController) RenderJson(c *gin.Context,data map[string]interface{},)  {
	c.JSON(200,gin.H(data))
}

func (bc *BaseController) HandlerError(c *gin.Context,message string,errInf string)  {
	c.JSON(500,gin.H{"message":message,"err":errInf})
}