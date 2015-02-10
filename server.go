package main

import (
	"flag"
	"fmt"
	"log"

	"net"
	"net/http"
	"net/rpc"

	"github.com/zeptonaut/golang-rpc-example/service"
)

var port = flag.Int("port", 9999, "Port on which to start the server.")

func main() {
	flag.Parse()

	server := new(service.MultiplicationService)
	rpc.Register(server)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	fmt.Printf("Serving on localhost:%d\n", *port)
	http.Serve(l, nil)
}
