package v1

import (
	"fmt"
	"ginDemo/entity"
	"net/http"
	"strconv"

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

	//此处需要可逆加密！
	//plainPasswd := passwd.Password
	//encrptPasswd, _ := utils.GeneratePassword(plainPasswd)
	//passwd.Password = encrptPasswd

	if passwd.Name != "" && passwd.Username != "" {
		flag, itemId := passwd.InsertPasswd()
		fmt.Println("item id: ", itemId)
		fmt.Println(flag)
		if !flag {
			fmt.Println("add passwd fail.")
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  "鉴权失败",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "add passwd success.",
	})
	return
}

func DelPasswd(c *gin.Context) {
	var passwd entity.Passwd
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	sqlId := passwd.DeletePassword(id).Id
	fmt.Println("sqlId: ", sqlId)
	//if !flag {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    -1,
	//		"message": "delete error",
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "delete success",
	})
}
