package main

import (
	"context"
	"fmt"
	"log"

	"gitub.com/zahrou/grpc/sum"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Couldn't connect %v", err)
	}

	defer conn.Close()

	summe := sum.NewSumServiceClient(conn)

	message := &sum.Variabls{Num: 12, Num2: 12}

	response, err := summe.SumThevalue(context.Background(), message)

	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't connect %v", err))
	}

	log.Printf("the response is : %d", response.Thesum)

}
