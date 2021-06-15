package v1

import (
	"encoding/json"
	"fmt"
	"ginDemo/common"
	"ginDemo/entity"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	var user entity.User
	//err := c.ShouldBind(&user)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 2001,
	//		"msg":  "无效的参数",
	//	})
	//	return
	//}
	var params map[string]string
	//根据body创建一个解析器
	decoder := json.NewDecoder(c.Request.Body)
	//解析参数，存入map
	decoder.Decode(&params)

	Username := params["username"]
	Password := params["password"]

	// 判断当前用户是否存在？如存在获取密码
	dbUser := user.Inquire(Username)
	//fmt.Println("dbUser pass: ", dbUser.Password)
	//fmt.Println("user pass: ", Password)
	if dbUser.Username != "" {
		// 获取并比较密码是否相同
		flag, err := utils.CheckPassword(dbUser.Password, Password)
		fmt.Println("flag: ", flag)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  "鉴权失败",
			})
			return
		}
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
