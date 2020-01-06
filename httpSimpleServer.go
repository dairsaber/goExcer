package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this is my first go server"))
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8001", nil)
}
