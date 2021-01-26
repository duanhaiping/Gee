package main

import (
	"base1"
	"net/http"
)

func main() {
	r := base1.New()
	r.GET("/", func(c *base1.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *base1.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.Run(":8080")

}
