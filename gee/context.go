package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Cont map[string]interface{}

type Context struct {
	//origin obj
	rsp http.ResponseWriter
	req *http.Request
	//Request info
	Path   string
	Method string
	Params map[string]string
	//Rsp info
	StatusCode int
}

func newContext(rsp http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		rsp:    rsp,
		req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
	c.rsp.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.rsp.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(code)
	c.rsp.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatusCode(code)
	encoder := json.NewEncoder(c.rsp)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.rsp, err.Error(), 500)
	}
}

func (c *Context) Data(code int, str []byte) {
	c.SetStatusCode(code)
	c.rsp.Write(str)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatusCode(code)
	c.rsp.Write([]byte(html))
}

func (c *Context) Query(key string) string {
	return c.req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.req.FormValue(key)
}
