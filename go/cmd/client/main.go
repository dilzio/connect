package main

import (
	"connectrpc.com/connect"
	"context"
	greetv1 "go_example/gen/greet/v1"
	"go_example/greetservice"
	//"go_example/gen/greetv1connect"
	"log"
	//"net/http"
)

func main() {
	//Call with Connect protocol
	//connectClient := greetv1connect.NewGreetServiceClient(http.DefaultClient, "http://localhost:8080")
	//res, err := connectClient.Greet(context.Background(), connect.NewRequest(&greetv1.GreetRequest{Name: "Jack"}))

	//Call with gRPC
	//gRPCClient := greetv1connect.NewGreetServiceClient(http.DefaultClient, "http://localhost:8080", connect.WithGRPC())
	//res, err := gRPCClient.Greet(context.Background(), connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}))

	//Call with gRPC Web
	//gRPCWebClient := greetv1connect.NewGreetServiceClient(http.DefaultClient, "http://localhost:8080", connect.WithGRPCWeb())
	//res, err := gRPCWebClient.Greet(context.Background(), connect.NewRequest(&greetv1.GreetRequest{Name: "Joe"}))

	//Call in process
	inProcess := &greetservice.GreetServer{}
	res, err := inProcess.Greet(context.Background(), connect.NewRequest(&greetv1.GreetRequest{Name: "John"}))

	if err != nil {
		log.Println(err)
	}
	log.Println(res.Msg.Greeting)
}