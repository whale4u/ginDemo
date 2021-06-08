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

	// v1 版本
	GroupV1 := r.Group("/v1")
	GroupV1.Use(jwt.JWTAuthMiddleware(), casbin.CasbinMiddleware())
	//GroupV1.Use(dummy.DummyMiddleware())
	{
		GroupV1.Any("/home/index", v1.Index)
	}
}
