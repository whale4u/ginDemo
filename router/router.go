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
	r.POST("/passwd/addpasswd", v1.AddPasswd)
	r.POST("/user/add", v1.AddUser)
	r.DELETE("/user/delete/:id", v1.DelUser)
	r.POST("/user/find", v1.FindUser)
	r.POST("/user/update", v1.UpdateUser)
	r.GET("/user/getall", v1.GetUsers)

	// v1 版本
	GroupV1 := r.Group("/v1")
	GroupV1.Use(jwt.JWTAuthMiddleware(), casbin.CasbinMiddleware())
	//GroupV1.Use(dummy.DummyMiddleware())
	{
		GroupV1.Any("/home/index", v1.Index)
	}
}
