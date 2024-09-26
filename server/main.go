package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	Message string
}

type Handler int

func (h *Handler) Ping(args *Args, reply *Args) error {
	fmt.Println("Received message: ", args.Message)

	switch args.Message {
	case "ping", "Ping", "PING":
		reply.Message = "pong"
	default:
		reply.Message = "I don't understand"
	}

	fmt.Println("Sending message: ", reply.Message)
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

	h := new(Handler)
	rpc.Register(h)
	fmt.Println("Listening on... ", conn.Addr())
	defer conn.Close()
	rpc.Accept(conn)
}
