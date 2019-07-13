package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	pb "gassu_music/batch-service/pb"
	"log"
)

func connection() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprint("localhost:8081")
	err := pb.RegisterAiringHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatal("handler registration error:", err)
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	if err := connection(); err != nil {
		log.Fatal("connection error", err)
	}
	//sampleなのでwithInsecure
	// conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatal("client connection error:", err)
	// }
	// defer conn.Close()
	// client := pb.NewAiringClient(conn)
	// message := &pb.GetAiringMessage{AiringId: "bkf0rlp94cg7ec81mi1g"}
	// res, err := client.GetAiring(context.TODO(), message)
	// fmt.Printf("result:%#v \n", res)
	// fmt.Printf("error::%#v \n", err)
}
