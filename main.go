package main

import (
	"Tycon/tycon"
	"fmt"
	"net/http"
)

func main() {
	r := tycon.New()
	r.GET("/", func(c *tycon.Context) {
		fmt.Println("1111")
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.Run(":9999")
}
