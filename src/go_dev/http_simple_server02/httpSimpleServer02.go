package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("============开始===========")
	fmt.Println("r.form===>", r.Form)
	fmt.Println("r.URL.Path===>", r.URL.Path)
	fmt.Println("r.URL.Scheme===>", r.URL.Scheme)
	for k,v := range r.Form {
		fmt.Println(k,v)
	}
	fmt.Println("=============结束==========")
	w.Write([]byte("hello,我已经收到你的请求"))
}
func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("heheda", handler)

	err := http.ListenAndServe("127.0.0.1:2020", nil)

	if err != nil {
		fmt.Println("监听端口失败",err)
	}
}
