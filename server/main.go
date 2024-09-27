package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	Message string
}

type Handler int

func (h *Handler) Ping(args *Args, reply *Args) error {
	log.Println("Received message:", args.Message)

	switch args.Message {
	case "ping", "Ping", "PING":
		reply.Message = "pong"
	default:
		reply.Message = "I don't understand"
	}

	log.Println("Sending message:", reply.Message)
	return nil
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:52648")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	h := new(Handler)
	log.Printf("Server listening at %v", conn.Addr())

	s := rpc.NewServer()
	s.Register(h)
	s.Accept(conn)
}
