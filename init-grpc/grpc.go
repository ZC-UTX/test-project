package main

import (
	"flag"
	"github.com/zchengutx/testproject/config"
	__ "github.com/zchengutx/testproject/topics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.TopicClient = __.NewTopicClient(conn)

}
