package gee

import (
	"net/http"
)

type HandleFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (engine *Engine) AddRoute(method string, pattern string, handleFunc HandleFunc) {
	engine.router.AddRoute(method, pattern, handleFunc)
}

func (engine *Engine) GET(pattern string, handleFunc HandleFunc) {
	engine.router.AddRoute("GET", pattern, handleFunc)
}

func (engine *Engine) POST(pattern string, handleFunc HandleFunc) {
	engine.router.AddRoute("POST", pattern, handleFunc)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	engine.router.handle(c)
}
