package main

import (
	"context"
	"log"

	"github.com/abhishekshree/gRPC-ws/chat"
	"google.golang.org/grpc"
)

func main() {
	PORT := ":9000"

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatClient(conn)

	msg := chat.Message{
		Body: "Hello World from the Client",
	}

	res, err := c.SendMessage(context.Background(), &msg)
	if err != nil {
		log.Fatalf("Error when calling SendMessage: %v", err)
	}
	log.Printf("Response from server: %v", res.Body)
}
