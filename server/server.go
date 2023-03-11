package main

import (
	"fmt"
	"log"
	"net"

	"gitub.com/zahrou/grpc/sum"
	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal("")
	}

	service := sum.Server{}
	grpcServer := grpc.NewServer()

	sum.RegisterSumServiceServer(grpcServer,&service)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(fmt.Sprintf("Failed to serve gRPC server over port 9000 %v", err))
	}
}
