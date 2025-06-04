package init_grpc

import (
	"flag"
	"github.com/zchengutx/testproject/config"
	__2 "github.com/zchengutx/testproject/works"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func InitWorks() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.WorksClient = __2.NewWorksClient(conn)

}
