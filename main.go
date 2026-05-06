package main

import (
	"fmt"
	"net/http"
)

type DefaultHandler struct{}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	path := r.URL.Path
	method := r.Method
	message := fmt.Sprintf(`Host: %s
Path: %s
Method: %s`, host, path, method)
	w.Write([]byte(message))
}

func main() {
	handler := &DefaultHandler{}
	http.ListenAndServe(":8080", handler)
}
