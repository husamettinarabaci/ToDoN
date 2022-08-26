package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var addTests = []string{
	"Wake up",
	"Prepare breakfast",
	"Get dressed",
	"Go to work",
}

func setServerConfigForTest() {
	memcacheServerIP = "a06cfcc7dfbca4c829963fcf485472bb-53359176.eu-central-1.elb.amazonaws.com"
}

func TestHealthHandler(t *testing.T) {
	t.Run("health probe statu", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		response := httptest.NewRecorder()

		HealthHandler(response, request)

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("Health Probe has an error : %v", response.Result().StatusCode)
		}
	})
}

func TestGetAllTodoHandler(t *testing.T) {
	setServerConfigForTest()
	t.Run("get all todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		GetAllTodoHandler(response, request)

		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("GetAllTodoHandler has an error : %v", response.Result().StatusCode)
		}
	})
}

func TestAddTodoHandler(t *testing.T) {
	setServerConfigForTest()
	t.Run("add todo", func(t *testing.T) {
		for _, v := range addTests {
			request, _ := http.NewRequest(http.MethodPost, "/add", strings.NewReader(fmt.Sprintf("item=%s", v)))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			response := httptest.NewRecorder()

			AddTodoHandler(response, request)

			if response.Result().StatusCode != http.StatusMovedPermanently {
				t.Errorf("AddTodoHandler has an error : %v", response.Result().StatusCode)
			}
		}
	})
}
