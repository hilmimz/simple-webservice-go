package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"simple-webservice/internal/models"
	"simple-webservice/internal/services"
	"simple-webservice/internal/test"
	"strings"
	"testing"
)

// Test Handler Average Uptime
func TestAvgUptime_Success(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 60},
						{Uptime: 120},
						{Uptime: 180},
						{Uptime: 240},
						{Uptime: 0},
						{Uptime: 0},
						{Uptime: 0},
						{Uptime: 0},
					},
				},
			}, nil
		},
	}
	routerService := services.NewRouterService(mockRepo)
	routerHandler := NewRouterHandler(routerService)

	req := httptest.NewRequest(http.MethodGet, "/api/uptime/avg", nil)
	rr := httptest.NewRecorder()

	routerHandler.AvgUptime(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	expected := `{"Router 1":75}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("unexpected body: %s", rr.Body.String())
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("content-type not set")
	}
}

func TestAvgUptime_Error(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return nil, errors.New("json file rror")
		},
	}

	routerService := services.NewRouterService(mockRepo)
	routerHandler := NewRouterHandler(routerService)

	req := httptest.NewRequest(http.MethodGet, "/api/uptime/avg", nil)
	rr := httptest.NewRecorder()

	routerHandler.AvgUptime(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "failed to process data") {
		t.Errorf("unexpected error message: %s", rr.Body.String())
	}
}

// End of Test Handler Average Uptime

// Test Handler Availability
func TestAvailability_Success(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 60},
						{Uptime: 120},
						{Uptime: 180},
						{Uptime: 240},
						{Uptime: 0},
						{Uptime: 0},
						{Uptime: 0},
						{Uptime: 0},
					},
				},
			}, nil
		},
	}
	routerService := services.NewRouterService(mockRepo)
	routerHandler := NewRouterHandler(routerService)

	req := httptest.NewRequest(http.MethodGet, "/api/uptime/availability", nil)
	rr := httptest.NewRecorder()

	routerHandler.Availability(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	expected := `{"Router 1":100}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("unexpected body: %s", rr.Body.String())
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("content-type not set")
	}
}

func TestAvailability_Error(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return nil, errors.New("json file rror")
		},
	}

	routerService := services.NewRouterService(mockRepo)
	routerHandler := NewRouterHandler(routerService)

	req := httptest.NewRequest(http.MethodGet, "/api/uptime/avg", nil)
	rr := httptest.NewRecorder()

	routerHandler.Availability(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "failed to process data") {
		t.Errorf("unexpected error message: %s", rr.Body.String())
	}
}
