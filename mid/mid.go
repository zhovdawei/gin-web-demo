package mid

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TestHandler(c *gin.Context) {
	fmt.Println("请求处理函数")
	c.JSON(http.StatusOK, gin.H{
		"message": "/test",
	})
}

func MidTest(c *gin.Context) {
	fmt.Println("mid tool start...")

	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("mid tool end...")
}
