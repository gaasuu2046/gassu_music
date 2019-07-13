package main

import (
	pb "gassu_music/batch-service/pb"
	"gassu_music/batch-service/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	airingService := &service.AiringService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterAiringServer(server, airingService)
	server.Serve(listenPort)
}
