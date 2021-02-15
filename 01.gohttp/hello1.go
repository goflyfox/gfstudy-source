package main

import (
	"fmt"
	"net/http"
)

//③处理请求，返回结果
func Hello1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world1")
}

func main() {
	//①路由注册
	http.HandleFunc("/", Hello1)
	//②服务监听
	http.ListenAndServe(":8080", nil)
}
