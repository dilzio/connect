package main

import (
	"go_example/gen/greet/v1/greetv1connect"
	"go_example/greetservice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func main() {
	greeter := &greetservice.GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
}
