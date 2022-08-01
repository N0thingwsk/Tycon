package main

import (
	"Tycon/tycon"
	"net/http"
)

func main() {
	r := tycon.New()
	r.GET("/", func(context *tycon.Context) {
		context.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *tycon.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *tycon.Context) {
		c.JSON(http.StatusOK, tycon.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
