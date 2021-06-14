package v1

import (
	"encoding/json"
	"fmt"
	"ginDemo/entity"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddUser(c *gin.Context) {
	var user entity.User
	var params map[string]string
	//根据body创建一个解析器
	decoder := json.NewDecoder(c.Request.Body)
	//解析参数，存入map
	decoder.Decode(&params)

	user.Username = params["username"]
	PlainPassword := params["password"]
	user.Password, _ = utils.GeneratePassword(PlainPassword)
	if user.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "username can't be empty ",
		})
		return
	}

	result := user.Inquire(user.Username)
	if result.Username != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "add user error",
		})
		return
	}
	_, err := user.Add()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "add user error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"data": "add user ok",
	})
}

func DelUser(c *gin.Context) {
	var user entity.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	_, err = user.Delete(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "delete error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "delete success",
	})
}
func FindUser(c *gin.Context) {
	var user entity.User
	username := c.Request.FormValue("username")
	//fmt.Println("username: ", username)
	result := user.Inquire(username)
	if result.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "find user error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": result,
	})
}

func UpdateUser(c *gin.Context) {
	var user entity.User
	var params map[string]string
	//根据body创建一个解析器
	decoder := json.NewDecoder(c.Request.Body)
	//解析参数，存入map
	decoder.Decode(&params)

	username := params["username"]
	PlainPassword := params["password"]
	EncryptPassword, _ := utils.GeneratePassword(PlainPassword)

	fmt.Println("username: ", username)
	result := user.Inquire(username)
	if result.Username != "" {
		fmt.Println("found user: ", result.Username)
		if user.Change(username, EncryptPassword) {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "update ok",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "update user error",
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "some error",
		})
		return
	}
}

func GetUsers(c *gin.Context) {
	var user entity.User
	users := user.Users()
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": users,
	})
	return
}
