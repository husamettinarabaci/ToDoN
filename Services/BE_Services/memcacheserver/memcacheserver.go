package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	itempb "proto/item"

	"google.golang.org/grpc"
)

type itemPbServer struct {
	itempb.UnimplementedSvcItemServer
}

var (
	errDublicateValue = errors.New("duplicate value is ignored")
	errEmptyValue     = errors.New("value can not be empty")
)

var (
	cachedItems itempb.PbItems
	chItem      chan string
)

var (
	serverPort string
)

var (
	wg sync.WaitGroup
)

var (
	isStarted = false
)

// init
// The value of serverPort is setting to :33800 for default
// If "SERVER_PORT" variable is exist in the environment
// then the value of serverPort is setting to ":SERVER_PORT" variable
func init() {
	serverPort = ":33800"
	if sp := os.Getenv("SERVER_PORT"); sp != "" {
		serverPort = ":" + sp
	}
	chItem = make(chan string)
}

func main() {

	wg.Add(3)
	StartApp()
	wg.Wait()
}

// StartApp starts application once
func StartApp() {
	if isStarted {
		return
	}
	isStarted = true
	go storeValues()
	go CreategRPCServer()
	go createHTTPServer()

	time.Sleep(time.Second * 5)
}

// createHTTPServer creates a new http-server and listens it
// This server is used for probe by K8S
func createHTTPServer() {
	http.HandleFunc("/health", HealthHandler)
	log.Println("Server listening at :80")
	http.ListenAndServe(":80", nil)
}

// HealthHandler is used by K8S for probe of healty
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// CreategRPCServer creates a new grpc-server and listens it
func CreategRPCServer() {
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

// RPCItem receives the new value and store it in the mem-cache
// This function returns a message of "SUCCESS" or "Error message"
func (s *itemPbServer) RPCItem(ctx context.Context, in *itempb.PbItem) (*itempb.PbResp, error) {
	var pbResp itempb.PbResp

	sleepASecond()

	//Check if the value is empty
	if in.Value == "" {
		pbResp.IsErr = true
		pbResp.Error = errEmptyValue.Error()
	} else {

		//Check if the received data exist in the cache
		isExist := false

		for _, v := range cachedItems.Items {
			if v.Value == in.Value {
				isExist = true
				break
			}
		}
		if !isExist {
			chItem <- in.Value
			pbResp.Message = "SUCCESS"
		} else {
			pbResp.IsErr = true
			pbResp.Error = errDublicateValue.Error()
		}
	}

	return &pbResp, nil
}

// RPCItems returns all values in the mem-cache
func (s *itemPbServer) RPCItems(ctx context.Context, in *itempb.PbReq) (*itempb.PbItems, error) {

	sleepASecond()
	return &cachedItems, nil
}

// sleepASecond blocks for a second the request for possible DDOS Attack
func sleepASecond() {

	time.Sleep(time.Second * 1)

}

// storeValues stores received values in to the mem-cache
func storeValues() {

	for newItem := range chItem {
		cachedItems.Items = append(cachedItems.Items, &itempb.PbItem{Value: newItem})
	}
}
