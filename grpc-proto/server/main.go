package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	um "shiva/proto/models/user"
	us "shiva/proto/services/user"

	"google.golang.org/grpc"
)

type UserApi struct {
	us.UnimplementedUserApiServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	us.RegisterUserApiServer(s, &UserApi{})

	fmt.Println("server started ...")
	createResponse()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server stopeed: %v", err)
	}
}

var (
	res *um.UserResponse
)

func createResponse() {
	res = &um.UserResponse{}
	b, _ := ioutil.ReadFile("op.json")
	json.Unmarshal(b, res)

}

func (*UserApi) AddUser(ctx context.Context, req *um.UserRequest) (*um.UserResponse, error) {
	return res, nil
}
