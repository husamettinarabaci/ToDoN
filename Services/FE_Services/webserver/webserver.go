package main

import (
	"context"
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
	http.HandleFunc("/add", AddTodoHandler)
	http.HandleFunc("/health", HealthHandler)
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
		Todos: todos,
	}

	_ = indexView.Execute(w, data)
}

type viewDatas struct {
	Todos []string
}
