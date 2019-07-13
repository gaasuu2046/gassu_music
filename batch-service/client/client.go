package main

import (
	"context"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "gassu_music/batch-service/pb"
	"log"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewAiringClient(conn)
	message := &pb.GetAiringMessage{AiringId: "bkf0rlp94cg7ec81mi1g"}
	res, err := client.GetAiring(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
