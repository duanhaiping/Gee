package GeeWeb

import (
	"log"
	"net/http"
	"strings"
)

/*
HandlerFunc

*/
type HandlerFunc func(c *Context)

type Engine struct {
	// 路由映射表
	router *router

	//将Engine作为最顶层的分组，也就是说Engine拥有RouterGroup所有的能力

	*RouterGroup
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
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
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

/*
添加路由
method 为GET POST PUT DELETE PATCHD等
为底层基础方法
*/
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Print("route %4s -%s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

/*
添加Get 请求路由
*/
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute(Method_Get, pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute(Method_POST, pattern, handler)
}
func (group *RouterGroup) DELETE(pattern string, handle HandlerFunc) {
	group.addRoute(Method_DELETE, pattern, handle)
}
func (group *RouterGroup) PUT(pattern string, handle HandlerFunc) {
	group.addRoute(Method_PUT, pattern, handle)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	engine.router.handle(c)
}
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
