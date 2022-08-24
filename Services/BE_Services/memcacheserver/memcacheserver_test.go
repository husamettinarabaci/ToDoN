package main

import (
	"net/http"
	"net/http/httptest"
	itempb "proto/item"
	"testing"
)

var addTests = []itempb.PbItem{
	itempb.PbItem{Value: "Wake up"},
	itempb.PbItem{Value: "Prepare breakfast"},
	itempb.PbItem{Value: "Get dressed"},
	itempb.PbItem{Value: "Go to work"},
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
