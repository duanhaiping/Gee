package GeeWeb

import (
	"net/http"
)

/*
HandlerFunc

*/
type HandlerFunc func(c *Context)

type Engine struct {
	// 路由映射表
	router *router
}

const (
	Method_Get    string = "GET"
	Method_POST   string = "POST"
	Method_PUT    string = "PUT"
	Method_DELETE string = "DELETE"
)

/*
创建实例
*/
func New() *Engine {
	return &Engine{router: newRouter()}
}

/*
添加路由
method 为GET POST PUT DELETE PATCHD等
为底层基础方法
*/
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

/*
添加Get 请求路由
*/
func (engine Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute(Method_Get, pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute(Method_POST, pattern, handler)
}
func (engine Engine) DELETE(pattern string, handle HandlerFunc) {
	engine.addRoute(Method_DELETE, pattern, handle)
}
func (engine Engine) PUT(pattern string, handle HandlerFunc) {
	engine.addRoute(Method_PUT, pattern, handle)
}

func (engine Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
