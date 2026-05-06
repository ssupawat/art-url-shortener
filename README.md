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
