package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := NewRouter()
	r.AddRoute("GET", "/", nil)
	r.AddRoute("GET", "/hello/:name", nil)
	r.AddRoute("GET", "/hello/b/c", nil)
	r.AddRoute("GET", "/hi/:name", nil)
	r.AddRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.GetRoute("GET", "/hello/yexiaoyu")
	n1, ps1 := r.GetRoute("GET", "/assets/css/yexiaoyu.css")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n1 == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if n1.pattern != "/assets/*filepath" {
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "yexiaoyu" {
		t.Fatal("name should be equal to 'yexiaoyu'")
	}
	if ps1["filepath"] != "css/yexiaoyu.css" {
		t.Fatal("name should be equal to 'css/yexiaoyu.css'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
	fmt.Printf("matched path: %s, params['name']: %s\n", n1.pattern, ps1["filepath"])
}
