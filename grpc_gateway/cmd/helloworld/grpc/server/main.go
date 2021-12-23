package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	helloworldpb "github.com/gojustforfun/learn-by-test/grpc_gateway/proto/helloworld"
	"github.com/gojustforfun/learn-by-test/grpc_gateway/server"
)

func main() {

	ln, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	
	helloworldpb.RegisterGreeterServer(s, server.NewHelloworld())

	// Serve gRPC Server
	log.Println("Serving gRPC on localhost:8080")
	log.Fatal(s.Serve(ln))
}
