# go-ecommerce-app

## How it works
This codebase eventually results in two different binaries; one for the web server, one for API server.

### Webserver
The web server is defined in the [./cmd/web](./cmd/web/main.go) directory. It's creates a new `application` struct
that wraps up the standard [http.Server](https://pkg.go.dev/net/http#Server), adds loggers, defines routes by providing
a handler and starts it.

It listens and accepts requests on http://localhost:4000.

## Resources
https://www.udemy.com/course/building-web-applications-with-go-intermediate-level/
