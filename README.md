# golang-rpc-example
Simple server and client that communicate using Go's RPC library.

Start the server with:

    > go run server.go

Then send an RPC from the client with:

    > go run client.go --factor1=12 --factor2=13
    156
