package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware() func(c *gin.Context) {
	//在return语句之前定义的内容将只在初始化时执行一次。
	fmt.Println("casbin init")
	e, _ := casbin.NewEnforcer("/Users/whale4u/Code/ginDemo/config/acl_model.conf", "/Users/whale4u/Code/ginDemo/config/acl_policy.csv")

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
		ok, _ := e.Enforce(role, c.URL.Path, c.Method)

		if ok {
			c.Next()
			fmt.Println("casbin check pass")
		} else {
			fmt.Println("casbin check fail")
		}

		c.Next()
	}
}

func KeyMatch(key1 string, key2 string) bool {
	return key1 == key2
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
