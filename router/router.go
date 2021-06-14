package router

import (
	"ginDemo/controller/v1"
	"ginDemo/middleware/casbin"
	"ginDemo/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	//r.GET("/sn", SignDemo)
	r.POST("/home/login", v1.Login)
	// 密码管理接口
	r.POST("/passwd/addpasswd", v1.AddPasswd)
	r.DELETE("/passwd/delete/:id", v1.DelPasswd)
	// 用户管理接口
	r.POST("/user/add", v1.AddUser)
	r.DELETE("/user/delete/:id", v1.DelUser)
	r.POST("/user/find", v1.FindUser)
	r.POST("/user/update", v1.UpdateUser)
	r.GET("/user/getall", v1.GetUsers)

	// v1 版本
	GroupV1 := r.Group("/v1")
	GroupV1.Use(jwt.JWTAuthMiddleware(), casbin.CasbinMiddleware())
	{
		GroupV1.Any("/home/index", v1.Index)
	}
}
