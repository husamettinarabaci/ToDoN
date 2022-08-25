package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*var addTests = []string{
	"Wake up",
	"Prepare breakfast",
	"Get dressed",
	"Go to work",
}*/

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

/*func TestGetAllTodoHandler(t *testing.T) {
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
}*/
