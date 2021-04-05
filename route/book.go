package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookGroup(r *gin.Engine) {
	bookGroup := r.Group("/book")
	{
		bookGroup.GET("/one", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "/book/one"})
		})
		bookGroup.GET("/two", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "/book/two"})
		})
		bookGroup.GET("/three", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "/book/three"})
		})
	}
}
