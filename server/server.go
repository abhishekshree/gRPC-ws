package main

import (
	"log"
	"net"

	"github.com/abhishekshree/gRPC-ws/chat"
	"google.golang.org/grpc"
)

func main() {
	PORT := ":9000"
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", PORT, err)
	}

	s := chat.Server{}

	server := grpc.NewServer()
	chat.RegisterChatServer(server, &s)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
