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
	memcacheServerIp   string
	memcacheServerPort string
)

// init
// The values of memcacheServerIp and memcacheServerPort are setting to localhost:33800 for default
// If "MEMCACHE_SERVER_IP" or "MEMCACHE_SERVER_PORT" variables are exist in the environment
// then the values of memcacheServerIp and memcacheServerPort are setting to "MEMCACHE_SERVER_IP:MEMCACHE_SERVER_PORT" variables
func init() {
	memcacheServerIp = "localhost"
	memcacheServerPort = "33800"
	if si := os.Getenv("MEMCACHE_SERVER_IP"); si != "" {
		memcacheServerIp = si
	}
	if sp := os.Getenv("MEMCACHE_SERVER_PORT"); sp != "" {
		memcacheServerPort = sp
	}
}

func main() {

	wg.Add(1)
	go createHttpServer()
	wg.Wait()
}

// createHttpServer
// This function creates a new http-server and listens it
func createHttpServer() {
	http.HandleFunc("/", GetAllTodoHandler)
	http.HandleFunc("/add", AddTodoHandler)
	http.HandleFunc("/health", HealthHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

// HealthHandler
// This handler is used by K8S for probe of healty
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// AddTodoHandler
// This handler adds a new todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	conn, err := grpc.Dial(memcacheServerIp+":"+memcacheServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := itempb.NewSvcItemClient(conn)
	resp, err := client.RpcItem(context.Background(), &itempb.PbItem{Value: item})
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

// GetAllTodoHandler
// This handler returns all todo values
func GetAllTodoHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(memcacheServerIp+":"+memcacheServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := itempb.NewSvcItemClient(conn)
	resp, err := client.RpcItems(context.Background(), &itempb.PbReq{})
	if err != nil {
		panic(err)
	}
	todos := make([]string, 0)

	if resp.Items != nil {
		for _, v := range resp.Items {
			todos = append(todos, v.Value)
		}
	}

	data := View{
		Todos: todos,
	}

	_ = indexView.Execute(w, data)
}

type View struct {
	Todos []string
}
