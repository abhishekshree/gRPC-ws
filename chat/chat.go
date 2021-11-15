package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SendMessage(ctx context.Context, msg *Message) (*Message, error) {
	log.Printf("Received message form client: %s", msg.Body)

	return &Message{Body: "Hello World from the Server"}, nil
}
