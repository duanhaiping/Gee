package main

import (
	"GeeWeb"
	"log"
	"net/http"
	"time"
)

func main() {
	r := GeeWeb.Default()

	r.Use(GeeWeb.Logger())
	r.GET("/", func(c *GeeWeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/panic", func(c *GeeWeb.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{

		v2.GET("/hello/:name", func(c *GeeWeb.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}
	//v2.Use(onlyForV2())
	/*v2.GET("/hello/:name", func(c *GeeWeb.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})*/

	r.Run(":8080")

}

func onlyForV2() GeeWeb.HandlerFunc {
	return func(c *GeeWeb.Context) {
		t := time.Now()
		c.Fail(http.StatusInternalServerError, "onlyFor2 :internal server error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
