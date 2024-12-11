package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"go_example/cmd/server/greetserver"
	greetv1 "go_example/gen"

	//"buf.build/gen/go/dilzio/bsrdemo/connectrpc/go/greetv1connect"
	"github.com/rs/cors"
	"go_example/gen/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type GreetServer struct{}

func (s *GreetServer) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
func main() {
	greeter := &greetserver.GreetServer{}
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
