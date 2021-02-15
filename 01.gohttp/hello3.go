package main

import (
	"fmt"
	"net/http"
)

type Hello3Handle struct {
}

//③处理请求，返回结果
func (handle *Hello3Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world32")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world31")
}

func main() {
	//①路由注册
	http.HandleFunc("/", Hello)
	//②服务监听
	http.ListenAndServe(":8080", &Hello3Handle{})
}
