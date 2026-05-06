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

func IsWhitelisted(ip string) bool {
	return slices.Contains(WhiteListIps, ip)
}

func IPWhitelistMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if !IsWhitelisted(ip) {
			http.Error(w, fmt.Sprintf("%s cannot access this http server", ip), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	user1 := &User{
		ID:   "1",
		Name: "Alice",
	}
	user2 := &User{
		ID:   "2",
		Name: "Bob",
	}

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "hello")
	})

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]*User{
			user1, user2,
		})
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		switch id {
		case "1":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user1)
		case "2":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user2)
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "not found")
		}
	})

	http.ListenAndServe(":8080", IPWhitelistMiddleware(mux))
}
