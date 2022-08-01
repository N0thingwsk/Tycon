package tycon

import (
	"net/http"
)

type HandlerFunc func(*Context)

// 路由表
type Engine struct {
	router *router
}

// 创建路由表实例
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// 添加路由
func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRouter(method, pattern, handler)
}

// GET GET请求方法
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

// POST POST请求方法
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// 解析路由表
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
