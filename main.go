package main

import "net/http"

type DefaultHandler struct{}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	handler := &DefaultHandler{}
	http.ListenAndServe(":8080", handler)
}
