# Note about this repo
**this README is completely written by human not AI**

First, I started this project while I don't remember about how to create http server in Go.

I know only how to do this:
1 go mod init
2 basic Go syntax
3 create the main.go with main func

And, I know that I will create the simplest version of url shorteners in my own way.

About the url shortner,
it is a system that contains:
1 POST /shorten body is {"long_url": "<long_url>"} returns {"short_url": "<short_url>"}
2 GET /<short_code> then redirect to the long url stored in a persistent module
3 GET /<short_code>/stats returns how many times this link was clicked

For how to do that let's research:

## How to create a HTTP server

I know that, in Go, the minimal way to create a HTTP server is using net/http which is stardard library from Go widely used by many web frameworks like Gin Gonic, Fiber, and Gorilla Mux.

After searching on the internet, I found that

there are 3 main parts of using it thanks for who that wrote this blog https://cyx.medium.com/using-the-net-http-package-in-go-fe219f6ab8c5

### Building a server
```
http.Handler
http.HandlerFunc
http.Server
```

### Writing a response
```
http.ResponseWriter
http.Request
```

### Making a request
```
http.Client
http.Transport
```

## What I learned
- Go type conversion: type(value) for example []byte("hello")
- extract a concrete type out of an interface by: value.(type) X cannot use with concrete types like 5.(int) this will error
- go run . vs go run main.go: go run main.go only reads main.go not other files but go run . reads all files in the current directory

## Version 1 - Hello net/http

```
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
```

the response body will be
```
hello
```

## Version 2 - More than 1 handler?
As you see, the HTTP server contain only 1 bahavior

responding 'hello'

How can we enhance it?

Read a request then route it to many response?
how can we separate differences between each request? -> path, method, body, header, remote address

### Other ways to do this
We need new handler that can help me do this

In Go or context of web servers, we have what can do this called Multiplexer or Mux

We can use http.ServeMux by directly register handle func though http.HandleFunc this will registor to http.DefaultServeMux or create new http.ServeMux by http.NewServeMux()

This is weighting between global and local mux

Ok, I think now we can move forward from net/http

## Version 3 - Get value from path

In Go 1.22+, we extract the id from the URL by using a named wildcard in our pattern and the r.PathValue() method in our handler.
https://go.dev/blog/routing-enhancements#:~:text=The%20new%20routing%20features%20almost,%E2%80%9D%20and%20has%20two%20segments.%20(
