package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:52648")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Please, inform the message:")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		args := Args{Message: scanner.Text()}
		log.Println("Sent message:", args.Message)
		reply := &Args{}
		err = client.Call("Handler.Ping", args, reply)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Received message:", reply.Message)
		log.Println("-------------------------")
	}
}
