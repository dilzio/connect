package greetservice

import (
	greetv1 "buf.build/gen/go/dilzio/bsrdemo/protocolbuffers/go"
	"connectrpc.com/connect"
	"context"
	"fmt"
	"log"
)

type TestGreetServer struct{}

func (s *TestGreetServer) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
