# go-ecommerce-app

## How it works
This codebase eventually results in two different binaries; one for the web server, one for API server.

### Webserver
The web server is defined in the [./cmd/web](./cmd/web/main.go) directory. It's creates a new `application` struct
that wraps up the standard [http.Server](https://pkg.go.dev/net/http#Server), adds loggers, defines routes by providing
a handler and starts it.

It listens and accepts requests on http://localhost:4000.

## Initial setup
* Set the `STRIPE_KEY` and `STRIPE_SECRET` environment variables first. For development purposes, you can create a
`.env` file in the root of this project and set the variables there. Check the Godotenv's
[Usage docs](https://github.com/joho/godotenv#usage) for details.

## Resources
https://www.udemy.com/course/building-web-applications-with-go-intermediate-level/
