package casbin

import (
	"fmt"
	//"reflect"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware() func(c *gin.Context) {
	//在return语句之前定义的内容将只在初始化时执行一次。
	fmt.Println("casbin init")
	e, _ := casbin.NewEnforcer("/Users/whale4u/Code/ginDemo/config/acl_model.conf", "/Users/whale4u/Code/ginDemo/config/acl_policy.csv")

	e.LoadPolicy()

	return func(c *gin.Context) {
		fmt.Println("in casbin")
		role, _ := c.Get("username")

		if role == "" {
			role = "anonymous"
		}
		// if it's a member, check if the user still exists
		//if role == "member" {
		//	uid, err := session.GetInt(r, "userID")
		//	if err != nil {
		//		writeError(http.StatusInternalServerError, "ERROR", w, err)
		//		return
		//	}
		//	exists := users.Exists(uid)
		//	if !exists {
		//		writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
		//		return
		//	}
		//}
		// casbin enforce
		//fmt.Println("====", reflect.TypeOf(role), reflect.TypeOf(c.Request.URL), reflect.TypeOf(c.Request.Method))
		ok, _ := e.Enforce(role, c.Request.URL.String(), c.Request.Method)
		//ok, _ := e.Enforce("admin111", "/v1/home/index", "GET")

		//获取所有policy
		//fmt.Println(e.GetPolicy())
		if ok {
			c.Next()
			fmt.Println("casbin check pass")
		} else {
			fmt.Println("casbin check fail")
			c.Abort()
			return
		}

		//c.Next()
	}
}
