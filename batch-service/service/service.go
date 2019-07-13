package service

import (
	"context"
	"fmt"
	pb "gassu_music/batch-service/pb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// AiringService struct
type AiringService struct {
}

var c = credentials.NewStaticCredentials("AKIA3FRSZ3IKDKLQWOPR", "FGb90o+x4bxSp2GH9/h2Ohfbfb66tLLWHN0u3Hvq", "") // 最後の引数は[セッショントークン]今回はなしで

// var db = dynamo.New(session.New(), &aws.Config{
// 	Credentials: c,
// 	Region:      aws.String("us-east-1"), // virginiaregion
// })
// var table = db.Table("Airing")

// GetAiring func
func (s *AiringService) GetAiring(ctx context.Context, message *pb.GetAiringMessage) (*pb.AiringResponse, error) {
	ddb := dynamodb.New(session.New(), &aws.Config{
		Credentials: c,
		Region:      aws.String("us-east-1"),
	})
	airingID := message.AiringId
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Airing"),

		Key: map[string]*dynamodb.AttributeValue{
			"airing_id": {
				S: aws.String(airingID),
			},
		},
	}

	resp, err := ddb.GetItem(params)

	if err != nil {
		fmt.Println(err.Error())
	}

	airing_json := fmt.Sprint((*resp).Item)
	return &pb.AiringResponse{
		// AiringJson: *resp.Item["airing_d"].S,
		AiringJson: airing_json,
	}, nil
}
