package main

import (
	"fmt"
	"net/http"
)

type Hello2Handle struct{}

//③处理请求，返回结果
func (handle *Hello2Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world2")
}

func main() {
	//①路由注册
	http.Handle("/", &Hello2Handle{})
	//②服务监听
	http.ListenAndServe(":8080", nil)
}
