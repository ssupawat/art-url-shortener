package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"slices"
)

var WhiteListIps = []string{"127.0.0.1"}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DefaultHandler struct{}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}

	if !slices.Contains(WhiteListIps, ip) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s cannot access this http server", ip)
		return
	}

	path := r.URL.Path
	method := r.Method
	if path == "/hello" && method == "GET" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "hello")
	} else if path == "/users" && method == "GET" {
		user := &User{
			ID:   "1234",
			Name: "Alice",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "not found")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "hello")
	})

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		user := &User{
			ID:   "1234",
			Name: "Alice",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	http.ListenAndServe(":8080", func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			if !slices.Contains(WhiteListIps, ip) {
				http.Error(w, fmt.Sprintf("%s cannot access this http server", ip), http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
		})
	}(mux))
}
