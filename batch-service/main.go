package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
)

type FileTransferToS3 struct {
	BucketName string
}

var c = credentials.NewStaticCredentials("AKIA3FRSZ3IKDKLQWOPR", "FGb90o+x4bxSp2GH9/h2Ohfbfb66tLLWHN0u3Hvq", "") // 最後の引数は[セッショントークン]今回はなしで

var svc = s3.New(session.New(), &aws.Config{
	Credentials: c,
	Region:      aws.String("us-east-1"),
})

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening 8000 ...")
	router.HandleFunc("/s3/list", GetList).Methods("GET")
	router.HandleFunc("/s3/upload/{supplier_id}", UploadFile).Methods("POST")
	log.Fatal(http.ListenAndServe(":8111", router))
}

// GetList S3
func GetList(w http.ResponseWriter, r *http.Request) {
	result, _ := svc.ListBuckets(nil)
	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
	json.NewEncoder(w).Encode(result)
	return
}

// UploadFile for Supplier to S3
func UploadFile(w http.ResponseWriter, r *http.Request) {
	result, _ := svc.ListBuckets(nil)
	fmt.Println("Buckets:")
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
