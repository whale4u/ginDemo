package dummy

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DummyMiddleware() func(c *gin.Context) {
	//在return语句之前定义的内容将只在初始化时执行一次。
	fmt.Println("execute only once!")

	return func(c *gin.Context) {
		fmt.Println("execute every time!")

		// Pass on to the next-in-chain

		c.Next()
	}
}
