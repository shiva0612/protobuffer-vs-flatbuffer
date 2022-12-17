package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	sm "shiv/api/models"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
)

var (
	req    *flatbuffers.Builder
	client sm.UserAPIClient
)

func getConnection() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure(), grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	return conn
}

func getrequest() *flatbuffers.Builder {
	var err error
	reqBytes, err := ioutil.ReadFile("ip.json")
	if err != nil {
		panic(err.Error())
	}

	UserRequestT := &sm.UserRequestT{}
	json.Unmarshal(reqBytes, UserRequestT)

	b := flatbuffers.NewBuilder(0)
	b.Finish(UserRequestT.Pack(b))

	return b

}

func init() {

	conn := getConnection()
	client = sm.NewUserAPIClient(conn)

	req = getrequest()

}
func BenchmarkFlatBuffer(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		res, err := client.Session_charging(ctx, req, grpc.CallContentSubtype("flatbuffers"))
		if err != nil {
			log.Println(err.Error())
		}
		_ = res
	}
}
