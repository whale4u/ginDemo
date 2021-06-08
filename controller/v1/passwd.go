package v1

import (
	"fmt"
	"ginDemo/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPasswd(c *gin.Context) {
	var passwd entity.Passwd
	err := c.ShouldBind(&passwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	fmt.Println(passwd)
	// 校验用户名和密码是否正确
	if passwd.UserName != "" {
		// 生成Token
		fmt.Println(passwd)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}
