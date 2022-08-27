package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"text/template"

	itempb "proto/item"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	wg sync.WaitGroup
)

var (
	indexView = template.Must(template.ParseFiles("./views/index.html"))
)

var (
	memcacheServerIP   string
	memcacheServerPort string
)

// init
// The values of memcacheServerIP and memcacheServerPort are setting to localhost:33800 for default
// If "MEMCACHE_SERVER_IP" or "MEMCACHE_SERVER_PORT" variables are exist in the environment
// then the values of memcacheServerIP and memcacheServerPort are setting to "MEMCACHE_SERVER_IP:MEMCACHE_SERVER_PORT" variables
func init() {
	memcacheServerIP = "localhost"
	memcacheServerPort = "33800"
	if si := os.Getenv("MEMCACHE_SERVER_IP"); si != "" {
		memcacheServerIP = si
	}
	if sp := os.Getenv("MEMCACHE_SERVER_PORT"); sp != "" {
		memcacheServerPort = sp
	}

	fmt.Println(memcacheServerIP)
}

func main() {

	wg.Add(1)
	go createHTTPServer()
	wg.Wait()
}

// createHTTPServer
// This function creates a new http-server and listens it
func createHTTPServer() {
	http.HandleFunc("/", GetAllTodoHandler)
	http.HandleFunc("/api/v1/all", GetAllTodoAPIHandler)
	http.HandleFunc("/add", AddTodoHandler)
	http.HandleFunc("/api/v1/add", AddTodoAPIHandler)
	http.HandleFunc("/health", HealthHandler)
	log.Println("Server listening at :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

// HealthHandler is used by K8S for probe of healty
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// AddTodoHandler adds a new todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	conn, err := grpc.Dial(memcacheServerIP+":"+memcacheServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := itempb.NewSvcItemClient(conn)
	resp, err := client.RPCItem(context.Background(), &itempb.PbItem{Value: item})
	if err != nil {
		panic(err)
	}
	if resp.Message == "SUCCESS" {
		log.Println("Todo value is added")
	} else if resp.IsErr {
		log.Printf("Todo value can't add : %v\n", resp.Error)
	} else {
		log.Println("RpcItem failed")
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// AddTodoAPIHandler adds a new todo by api
func AddTodoAPIHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var reqData addData
	json.Unmarshal(reqBody, &reqData)

	conn, err := grpc.Dial(memcacheServerIP+":"+memcacheServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := itempb.NewSvcItemClient(conn)
	resp, err := client.RPCItem(context.Background(), &itempb.PbItem{Value: reqData.Item})
	if err != nil {
		panic(err)
	}

	var message string = ""
	if resp.Message == "SUCCESS" {
		message = "OK"
		log.Println("Todo value is added")
	} else if resp.IsErr {
		message = "OK"
		log.Printf("Todo value can't add : %v\n", resp.Error)
	} else {
		message = "FAIL"
		log.Println("RpcItem failed")
	}

	data := viewDatas{
		Todos:   make([]string, 0),
		Message: message,
	}

	json.NewEncoder(w).Encode(data)
}

// GetAllTodoHandler returns all todo values
func GetAllTodoHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(memcacheServerIP+":"+memcacheServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := itempb.NewSvcItemClient(conn)
	resp, err := client.RPCItems(context.Background(), &itempb.PbReq{})
	if err != nil {
		panic(err)
	}
	todos := make([]string, 0)

	if resp.Items != nil {
		for _, v := range resp.Items {
			todos = append(todos, v.Value)
		}
	}

	data := viewDatas{
		Todos:   todos,
		Message: "",
	}

	_ = indexView.Execute(w, data)
}

// GetAllTodoAPIHandler returns all todo values by api
func GetAllTodoAPIHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(memcacheServerIP+":"+memcacheServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := itempb.NewSvcItemClient(conn)
	resp, err := client.RPCItems(context.Background(), &itempb.PbReq{})
	if err != nil {
		panic(err)
	}
	todos := make([]string, 0)

	if resp.Items != nil {
		for _, v := range resp.Items {
			todos = append(todos, v.Value)
		}
	}

	data := viewDatas{
		Todos:   todos,
		Message: "",
	}

	json.NewEncoder(w).Encode(data)
}

type viewDatas struct {
	Todos   []string
	Message string
}

type addData struct {
	Item string
}
