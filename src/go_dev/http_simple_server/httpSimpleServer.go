package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, I am a simple server by golang!"))
}

func main() {
	fmt.Println("服务启动在:127.0.0.1:8001地址上")
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8001", nil)

}
