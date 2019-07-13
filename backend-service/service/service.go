package service

import (
	"context"
	"fmt"
	pb "gassu_music/backend-service/pb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// AiringService struct
type AiringService struct {
}

var c = credentials.NewStaticCredentials("AKIA3FRSZ3IKDKLQWOPR", "FGb90o+x4bxSp2GH9/h2Ohfbfb66tLLWHN0u3Hvq", "") // 最後の引数は[セッショントークン]今回はなしで

var db = dynamo.New(session.New(), &aws.Config{
	Credentials: c,
	Region:      aws.String("us-east-1"), // virginiaregion
})
var table = db.Table("Airing")

// GetAiring func
func (s *AiringService) GetAiring(ctx context.Context, message *pb.GetAiringMessage) (*pb.AiringResponse, error) {

	airingID := message.AiringId

	var airing string
	err := table.Get("airing_id", airingID).All(&airing)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}
	return &pb.AiringResponse{
		// AiringJson: airing,
		AiringJson: airingID,
	}, nil
}
