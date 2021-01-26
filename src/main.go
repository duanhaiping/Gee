package main

import (
	"base1"
	"fmt"
	"net/http"
)

func main()  {
	r:=base1.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"url.path = %S \n",request.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":8081")

}