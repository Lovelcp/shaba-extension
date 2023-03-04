package main

import (
	"fmt"
	"net/http"
)

var value string

func main() {
	// 第一个HTTP GET方法，用于接收参数
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		// 从URL参数中获取值
		value = r.URL.Query().Get("value")

		// 将值打印到控制台
		fmt.Println(value)

		// 返回响应
		w.Write([]byte("Value received: " + value))
	})

	// 第一个HTTP GET方法，用于接收参数
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		// 返回响应
		w.Write([]byte(value))
	})

	// 启动HTTP服务器并监听22333端口
	http.ListenAndServe(":22333", nil)
}
