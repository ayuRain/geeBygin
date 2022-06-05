package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandleFunc
}

func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandleFunc),
	}
}

func parsePattern(pattern string) []string {
	PartOfRoute := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range PartOfRoute {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) AddRoute(method string, pattern string, handler HandleFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) GetRoute(method string, pattern string) (*node, map[string]string) {
	SearchParts := parsePattern(pattern)
	param := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.Search(SearchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, item := range parts {
			if item[0] == ':' {
				param[item[1:]] = SearchParts[index]
			}
			if item[0] == '*' && len(item) > 1 {
				param[item[1:]] = strings.Join(SearchParts[index:], "/")
				break
			}
		}
		return n, param
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, param := r.GetRoute(c.Method, c.Path)
	if n != nil {
		c.Params = param
		key := c.Method + "-" + n.pattern
		fmt.Println("%s %s", c.Path, n.pattern)
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
