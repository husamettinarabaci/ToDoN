package main

import (
	"context"
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

func TestHealthProbe(t *testing.T) {
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
	t.Run("add todo", func(t *testing.T) {

		ctx := context.Background()
		conn, err := grpc.Dial("localhost:33800", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			t.Errorf("Failed to dial server: %v", err)
		}
		defer conn.Close()
		client := itempb.NewSvcItemClient(conn)
		for _, v := range addTests {
			resp, err := client.RpcItem(ctx, &v)
			if err != nil {
				t.Errorf("RpcItem failed: %v", err)
			}
			if resp.Message == "SUCCESS" {
				t.Log("Todo value is added")
			} else if resp.IsErr == true {
				t.Logf("Todo value can't added : %v", resp.Error)
			} else {
				t.Error("RpcItem failed")
			}
		}
	})
}

func TestGetTodoList(t *testing.T) {
	t.Run("get all todo", func(t *testing.T) {

		ctx := context.Background()
		conn, err := grpc.Dial("localhost:33800", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			t.Errorf("Failed to dial server: %v", err)
		}
		defer conn.Close()
		client := itempb.NewSvcItemClient(conn)
		resp, err := client.RpcItems(ctx, &itempb.PbReq{})
		if err != nil {
			t.Errorf("RpcItems failed: %v", err)
		}
		if resp.Items == nil {
			t.Error("RpcItem failed: nil")
		}
		t.Logf("Returned values : %v", resp.Items)
	})
}
