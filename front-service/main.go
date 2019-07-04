package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/rs/xid"
)

// Airing Airing struct
type Airing struct {
	AiringID    xid.ID    `dynamo:"airing_id"`
	ProgramID   xid.ID    `dynamo:"program_id"`
	ProgramName string    `dynamo:"program_name"`
	CreatedTime time.Time `dynamo:"created_time"`
}

var c = credentials.NewStaticCredentials("AKIA3FRSZ3IKDKLQWOPR", "FGb90o+x4bxSp2GH9/h2Ohfbfb66tLLWHN0u3Hvq", "") // 最後の引数は[セッショントークン]今回はなしで

var db = dynamo.New(session.New(), &aws.Config{
	Credentials: c,
	Region:      aws.String("us-east-1"), // virginiaregion
})
var table = db.Table("Airing")

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening 8000 ...")
	router.HandleFunc("/airings", GetAirings).Methods("GET")
	router.HandleFunc("/airings/{id}", GetAiring).Methods("GET")
	router.HandleFunc("/airings", CreateAiring).Methods("POST")
	router.HandleFunc("/airings/{id}", UpdateAiring).Methods("PUT")
	router.HandleFunc("/airings/{id}", DeleteAiring).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// CreateAiring func　airingの作成
func CreateAiring(w http.ResponseWriter, r *http.Request) {
	programname := r.URL.Query().Get("name")
	airingid := xid.New()
	programid := xid.New()

	u := Airing{AiringID: airingid, ProgramID: programid, ProgramName: programname, CreatedTime: time.Now().UTC()}
	fmt.Println(u)

	if err := table.Put(u).Run(); err != nil {
		fmt.Println("err")
		panic(err.Error())
	}
}

// GetAirings func
func GetAirings(w http.ResponseWriter, r *http.Request) {
	var airings []Airing
	err := table.Scan().All(&airings)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	for i := range airings {
		fmt.Println(airings[i])
	}

	json.NewEncoder(w).Encode(airings)
	return
}

// GetAiring func
func GetAiring(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	airingID := params["id"]

	var airing []Airing
	err := table.Get("airing_id", airingID).All(&airing)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}
	fmt.Println(airing)
	json.NewEncoder(w).Encode(airing)
}

// UpdateAiring func
func UpdateAiring(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	q := r.URL.Query().Get("name")
	airingID := params["id"]
	err := table.Update("airing_id", airingID).Set("program_name", q).Run()
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}
}

// DeleteAiring func　id指定したユーザー情報を削除
func DeleteAiring(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	airingID := params["id"]
	table.Delete("airing_id", airingID).Run()
}
