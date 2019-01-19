package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	data, _ := c.GetRawData()
	result := CheckUser(data)
	c.JSON(200, result)
}

func UserInfo(c *gin.Context) {
	c.JSON(200, gin.H{"code": 20000,
		"data": map[string]interface{}{
			"roles":  []string{"admin"},
			"name":   "admin",
			"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"},
	})
}

func LogOut(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(500, gin.H{
			"code":    20000,
			"message": string(data),
		})
	}

	fmt.Println(string(data))
	c.JSON(200, gin.H{
		"code": 20000,
		"data": "success",
	})
}
