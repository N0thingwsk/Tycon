package tycon

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

// Engine 路由表
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// RouterGroup 路由分组
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

// New 创建路由表实例
func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	engine.groups = []*RouterGroup{
		engine.RouterGroup,
	}
	return engine
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		parent: r,
		engine: engine,
	}
	return newGroup
}

// 添加路由
func (r *RouterGroup) addRouter(method string, comp string, handler HandlerFunc) {
	pattern := r.prefix + comp
	log.Printf("Router %4s - %s", method, pattern)
	r.engine.router.addRouter(method, pattern, handler)
}

// GET GET请求方法
func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRouter("GET", pattern, handler)
}

// POST POST请求方法
func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRouter("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// 解析路由表
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
