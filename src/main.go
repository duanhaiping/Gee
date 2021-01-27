package main

import (
	"GeeWeb"
	"net/http"
)

func main() {
	r := GeeWeb.New()
	r.GET("/", func(c *GeeWeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *GeeWeb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
