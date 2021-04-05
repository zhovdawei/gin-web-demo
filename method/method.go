package method

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	// 方法1： 使用map
	/*data := map[string]interface{}{
		"name":    "小丸子",
		"message": "hello golang!",
		"age":     19,
	}*/

	// 方法2：gin.H
	data := gin.H{
		"name":    "小丸子",
		"message": "hello golang!",
		"age":     20}

	c.JSON(http.StatusOK, data)
}

func Hi(c *gin.Context) {
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}

	data := msg{
		"拜振华",
		"Hello pubs",
		29,
	}

	c.JSON(http.StatusOK, data)

}

func GetQueryParam(c *gin.Context) {
	//name := c.Query("query")
	//name := c.DefaultQuery("query","xxxxxx")
	name, ok := c.GetQuery("query")
	if !ok {
		name = "gggg"
	}

	c.JSON(http.StatusOK, gin.H{"name": name})
}

func GetUrlParam(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})

}

type UserInfo struct {
	Username string `form:"username" json:"account"`
	Password string `form:"password" json:"pwd"`
}

func Login(c *gin.Context) {

	var message string

	username := c.Query("username")
	if username == "" {
		message = "username 不能为空"
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
		return
	}

	password := c.Query("password")
	if password == "" {
		message = "password 不能为空"
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
		return
	}

	u := UserInfo{
		Username: username,
		Password: password,
	}
	message = "login success"
	fmt.Printf("%#v\n", u)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func LoginGet(c *gin.Context) {

	var message string

	var u UserInfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	message = "login success"
	fmt.Printf("%#v\n", u)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func LoginPostForm(c *gin.Context) {
	var message string

	var u UserInfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	message = "login success"
	fmt.Printf("%#v\n", u)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func LoginPostJson(c *gin.Context) {

	var message string

	var u UserInfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	message = "login success!!!!"
	fmt.Printf("%#v\n", u)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func DoBusiness(c *gin.Context) {
	var u UserInfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if u != (UserInfo{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "do business success",
		})
	}

	// 重定向
	c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/user")
}
