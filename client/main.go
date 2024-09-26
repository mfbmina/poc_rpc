package main

import (
	"bufio"
	"fmt"
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
		fmt.Println("Enter text: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		args := Args{Message: scanner.Text()}
		fmt.Println("Sent message: ", args.Message)
		reply := &Args{}
		err = client.Call("Handler.Ping", args, reply)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Received message: ", reply.Message)
		fmt.Println("-------------------------")
	}
}
