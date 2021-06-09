package v1

import (
	"ginDemo/common"
	"ginDemo/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	if user.Username == "admin" && user.Password == "pass" {
		// 生成Token
		tokenString, _ := common.GenToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

func Index(c *gin.Context) {
	val, _ := c.Get("username")
	name := val.(string)
	c.JSON(200, gin.H{
		"message": "home page",
		"user":    name,
	})
}
