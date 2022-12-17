package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	um "shiva/proto/models/user"
	us "shiva/proto/services/user"
	"testing"

	"google.golang.org/grpc"
)

var (
	client us.UserApiClient
	req    *um.UserRequest
)

func init() {

	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	client = us.NewUserApiClient(cc)

	req = &um.UserRequest{}

	b, _ := ioutil.ReadFile("ip.json")
	json.Unmarshal(b, req)

}

func unary() {
}

func BenchmarkProtoBuffer(b *testing.B) {

	b.ReportAllocs()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		res, err := client.AddUser(ctx, req)
		if err != nil {
			log.Println("error in client-unary: ", err)
		}
		_ = res
	}
}
