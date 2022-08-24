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
	cachedItems map[int]itempb.PbItem
	chItem      chan itempb.PbItem
	lastIndex   int
)

var (
	serverPort string
)

var (
	wg sync.WaitGroup
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

	chItem = make(chan itempb.PbItem)
	cachedItems = make(map[int]itempb.PbItem)

	wg.Add(3)
	go storeValues()
	go creategRpcServer()
	go createHttpServer()
	wg.Wait()
}

// createHttpServer
// This function creates a new http-server and listens it
// This server is used for probe by K8S
func createHttpServer() {
	http.HandleFunc("/health", HealthHandler)
	http.ListenAndServe(":80", nil)
}

// healthHandler
// This handler is used by K8S for probe of healty
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// creategRpcServer
// This function creates a new grpc-server and listens it
// It uses "SERVER_PORT" variable in the environment for the port of new server
func creategRpcServer() {
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

	log.Println("RpcItem : ")
	log.Println(in)

	sleepASecond()

	//Check if the value is empty
	if in.Value == "" {
		pbResp.IsErr = true
		pbResp.Error = errEmptyValue.Error()
	} else {

		//Check if the received data exist in the cache
		isExist := false

		for _, v := range cachedItems {
			if v.Value == in.Value {
				isExist = true
				break
			}
		}
		if !isExist {
			chItem <- *in
			pbResp.Message = "SUCCESS"
		} else {
			pbResp.IsErr = true
			pbResp.Error = errDublicateValue.Error()
		}
	}

	log.Println(pbResp)
	return &pbResp, nil
}

// RpcItems
// This function returns all values in the mem-cache
func (s *itemPbServer) RpcItems(ctx context.Context, in *itempb.PbReq) (*itempb.PbItems, error) {

	sleepASecond()
	items := &itempb.PbItems{}
	items.Items = make([]*itempb.PbItem, len(cachedItems))
	for i, v := range cachedItems {
		items.Items[i] = &itempb.PbItem{Value: v.Value}
	}
	log.Println("RpcItems : ")
	log.Println(cachedItems)
	log.Println(items)
	return items, nil
}

// sleepASecond
// This function blocks for a second the request for possible DDOS Attack
func sleepASecond() {

	time.Sleep(time.Second * 1)

}

// storeValues
// This function stores received values in to the mem-cache
func storeValues() {

	for newItem := range chItem {
		cachedItems[lastIndex] = newItem
		lastIndex++
		log.Println("Stored Values : ")
		log.Println(cachedItems)
	}
}
