package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/DataDavD/waytogocourse/ch13webapps/rpcobjects"
)

const serverAddress = "localhost"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &rpcobjects.Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)

	// Asynchronous call
	call1 := client.Go("Args.Multiply", args, &reply, nil)
	replyCall := <-call1.Done
	fmt.Println(replyCall.Args)
}
