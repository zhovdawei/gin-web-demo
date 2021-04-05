package route

import (
	"gin-web-demo/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProductGroup(r *gin.Engine) {
	bookGroup := r.Group("/product")
	{
		bookGroup.POST("/iphone", func(c *gin.Context) {
			iphone := orm.Product{
				ID:    1,
				Name:  "苹果18代全透明手机",
				Color: "太空灰",
				Price: 289.99,
			}
			iphone.Insert()
			c.JSON(http.StatusOK, gin.H{"message": "insert new product success"})
		})
		bookGroup.DELETE("/iphone", func(c *gin.Context) {
			iphone := orm.Product{
				ID:    1,
				Name:  "苹果18代全透明手机",
				Color: "太空灰",
				Price: 289.99,
			}
			iphone.Delete()
			c.JSON(http.StatusOK, gin.H{"message": "insert new product success"})
		})
		bookGroup.PUT("/iphone", func(c *gin.Context) {
			iphone := orm.Product{
				ID:    1,
				Name:  "苹果18代全透明手机",
				Color: "太空灰",
				Price: 289.99,
			}
			iphone.Update()
			c.JSON(http.StatusOK, gin.H{"message": "insert new product success"})
		})
		bookGroup.GET("/iphone", func(c *gin.Context) {
			var iphone orm.Product
			iphone.Selete()
			c.JSON(http.StatusOK, iphone)
		})
	}
}
