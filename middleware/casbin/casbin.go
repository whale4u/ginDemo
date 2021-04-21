package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware() func(c *gin.Context) {
	//在return语句之前定义的内容将只在初始化时执行一次。
	fmt.Println("casbin init")
	e, _ := casbin.NewEnforcer("/Users/whale4u/Code/ginDemo/config/rbac_model.conf", "/Users/whale4u/Code/ginDemo/config/rbac_policy.csv")

	// 注册自定义函数
	e.AddFunction("my_func", KeyMatchFunc)

	// Load the policy from DB.
	e.LoadPolicy()

	sub := "susan" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	return func(c *gin.Context) {
		fmt.Println("in casbin")

		// Pass on to the next-in-chain
		ok, err := e.Enforce(sub, obj, act)

		// fmt.Println(ok)

		if err != nil {
			// handle err
			fmt.Printf("%s", err)
		}

		if ok == true {
			// permit alice to read data1
			fmt.Println(sub, obj, act, " Pass")
		} else {
			// deny the request, show an error
			fmt.Println(sub, obj, act, " Fail")
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
