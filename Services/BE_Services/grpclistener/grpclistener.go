package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	itempb "proto/item"

	"google.golang.org/grpc"
)

type itemPbServer struct {
	itempb.UnimplementedSvcItemServer
}

func main() {
	CreateServer()
}

// TODO
func CreateServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	itempb.RegisterSvcItemServer(s, &itemPbServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

// TODO
func (s *itemPbServer) RpcItem(ctx context.Context, in *itempb.PbItem) (*itempb.PbResp, error) {
	return &itempb.PbResp{}, nil
}

// TODO
func (s *itemPbServer) RpcItems(ctx context.Context, in *itempb.PbReq) (*itempb.PbItems, error) {
	return &itempb.PbItems{}, nil
}
