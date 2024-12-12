package main

import (
	"buf.build/gen/go/dilzio/bsrdemo/connectrpc/go/greetv1connect"
	"github.com/rs/cors"
	"go_example/greetservice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func main() {
	greeter := &greetservice.GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Grpc-Status", "Grpc-Message", "Grpc-Status-Details-Bin"},
	})
	mux.Handle(path, c.Handler(handler))
	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
}
