package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"time"

	itempb "proto/item"

	"google.golang.org/grpc"
)

type itemPbServer struct {
	itempb.UnimplementedSvcItemServer
}

var (
	errDublicateValue = errors.New("Duplicate value is ignored")
)

var (
	cachedItems itempb.PbItems
	lastItem    itempb.PbItem
	chItem      chan itempb.PbItem
)

var (
	serverPort string
)

// init
// The value of serverPort is setting to 33800 for default
// If "SERVER_PORT" variable is exist in the environment
// then the value of serverPort is setting to "SERVER_PORT" variable
func init() {
	serverPort = ":33800"
	if sp := os.Getenv("SERVER_PORT"); sp != "" {
		serverPort = ":" + sp
	}
}

func main() {
	go storeValues()
	createServer()
}

// createServer
// This function creates a new grpc-server and listens it
// It uses "SERVER_PORT" variable in the environment for the port of new server
func createServer() {
	lis, err := net.Listen("tcp", serverPort)
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

// RpcItem
// This function receives the new value and store it in the mem-cache
// This function returns a message of "SUCCESS" or "Error message"
func (s *itemPbServer) RpcItem(ctx context.Context, in *itempb.PbItem) (*itempb.PbResp, error) {

	var pbResp itempb.PbResp

	sleepASecond()

	//Check if the same data overlaps
	if lastItem.Value != in.Value {
		lastItem = *in
		chItem <- *in
		pbResp.Message = "SUCCESS"
	} else {
		pbResp.IsErr = true
		pbResp.Error = errDublicateValue.Error()
	}
	return &pbResp, nil
}

// RpcItems
// This function returns all values in the mem-cache
func (s *itemPbServer) RpcItems(ctx context.Context, in *itempb.PbReq) (*itempb.PbItems, error) {

	sleepASecond()

	return &cachedItems, nil
}

// sleepASecond
// This function blocks for a second the request for possible DDOS Attack
func sleepASecond() {

	time.Sleep(time.Second * 1)

}

// storeValues
// This function stores received values in to the mem-cache
func storeValues() {

	for {
		select {
		case newItem := <-chItem:
			cachedItems.Items = append(cachedItems.Items, &newItem)
		}
	}
}
