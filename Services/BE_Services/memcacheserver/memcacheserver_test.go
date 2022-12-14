package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	itempb "proto/item"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addTests = []itempb.PbItem{
	{Value: "Wake up"},
	{Value: "Prepare breakfast"},
	{Value: "Get dressed"},
	{Value: "Go to work"},
}

func TestHealthHandler(t *testing.T) {
	StartApp()
	t.Run("health probe statu", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		response := httptest.NewRecorder()

		HealthHandler(response, request)

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("Health Probe has an error : %v", response.Result().StatusCode)
		}
	})
}

func TestAddTodo(t *testing.T) {
	StartApp()
	t.Run("add todo", func(t *testing.T) {
		ctx := context.Background()
		conn, err := grpc.Dial(serverPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := itempb.NewSvcItemClient(conn)
		fmt.Println(serverPort)

		for i := 0; i < len(addTests); i++ {

			v := addTests[i].Value
			resp, err := client.RPCItem(ctx, &itempb.PbItem{Value: v})
			if err != nil {
				t.Errorf("RpcItem failed: %v", err)
			}
			if resp.Message == "SUCCESS" {
				t.Log("Todo value is added")
			} else if resp.IsErr {
				t.Logf("Todo value can't add : %v", resp.Error)
			} else {
				t.Error("RpcItem failed")
			}
		}
	})
}

func TestGetTodoList(t *testing.T) {
	t.Run("get all todo", func(t *testing.T) {
		ctx := context.Background()
		conn, err := grpc.Dial(serverPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := itempb.NewSvcItemClient(conn)
		resp, err := client.RPCItems(ctx, &itempb.PbReq{})
		if err != nil {
			t.Errorf("RpcItems failed: %v", err)
		}
		if resp.Items == nil {
			t.Error("RpcItem failed: nil")
		}
		t.Logf("Returned values : %v", resp.Items)
	})

}
