package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"Hello , Golang!!!",
		})
	})
	r.Run(":9090")
}
*/

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Http server start failed, err:%", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./static/template/hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	name := "小伙子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}
