package main

import (
	"gin-web-demo/method"
	"gin-web-demo/mid"
	"gin-web-demo/orm"
	"gin-web-demo/route"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		method.Hello(c)
	})

	r.GET("/hi", func(c *gin.Context) {
		method.Hi(c)
	})

	r.GET("/web", func(c *gin.Context) {
		method.GetQueryParam(c)
	})

	r.GET("/xx/:name/:age", func(c *gin.Context) {
		method.GetUrlParam(c)
	})

	r.GET("/user", func(c *gin.Context) {
		method.Login(c)
	})

	r.GET("/userGet", func(c *gin.Context) {
		method.LoginGet(c)
	})

	r.POST("/userPostForm", func(c *gin.Context) {
		method.LoginPostForm(c)
	})

	r.POST("/userPostJson", func(c *gin.Context) {
		method.LoginPostJson(c)
	})

	r.GET("/doBusiness", func(c *gin.Context) {
		method.DoBusiness(c)
	})

	r.GET("/a", func(c *gin.Context) {
		// 转发
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b method handle....",
		})
	})

	r.GET("/abc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
			"path":   "abc",
		})
	})
	r.POST("/abc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
			"path":   "abc",
		})
	})
	r.PUT("/abc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
			"path":   "abc",
		})
	})
	r.DELETE("/abc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
			"path":   "abc",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "404"})
	})

	// 分组路由
	route.BookGroup(r)
	route.ProductGroup(r)

	r.Use(mid.MidTest)

	// 中间件，生命周期钩子函数
	r.GET("/test", mid.TestHandler)

	// **************************************************************************
	// 1.连接mysql数据库
	db, err := gorm.Open("mysql", "root:David2021_@(192.168.91.102:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 2.表注册
	orm.TableRegistry(db)

	r.Run(":9090")
}
