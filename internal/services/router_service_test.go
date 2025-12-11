package services

import (
	"errors"
	"simple-webservice/internal/models"
	"simple-webservice/internal/test"
	"testing"
)

// Average Uptime Test

func TestAvgUptimeOK(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 0},
						{Uptime: 60},
						{Uptime: 120},
						{Uptime: 180},
						{Uptime: 0},
						{Uptime: 60},
					},
				},
			}, nil
		},
	}
	routerService := NewRouterService(mockRepo)

	result, err := routerService.AvgUptime()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := (60 + 120 + 180 + 60) / 6
	if resultData := result["Router 1"]; resultData != float32(expected) {
		t.Errorf("expected %v, got %v", expected, resultData)
	}
}

func TestAvgUptime_LoadError(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return nil, errors.New("load error")
		},
	}

	routerService := NewRouterService(mockRepo)

	_, err := routerService.AvgUptime()

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// End of Average Uptime Test

// Availability Test

func TestAvailability_LoadError(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return nil, errors.New("load error")
		},
	}

	routerService := NewRouterService(mockRepo)

	_, err := routerService.Availability()

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestAvailability_AllZero(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 0},
						{Uptime: 0},
						{Uptime: 0},
						{Uptime: 0},
					},
				},
			}, nil
		},
	}

	expected := 0

	service := NewRouterService(mockRepo)
	result, err := service.Availability()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data := result["Router 1"]; data != float32(expected) {
		t.Errorf("expected %v, got %v", expected, data)
	}
}

func TestAvailability_AllUp(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 60},
						{Uptime: 120},
						{Uptime: 180},
					},
				},
			}, nil
		},
	}

	expected := 100

	service := NewRouterService(mockRepo)
	result, err := service.Availability()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data := result["Router 1"]; data != float32(expected) {
		t.Errorf("expected %v, got %v", expected, data)
	}
}

func TestAvailability_PF(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 0},
						{Uptime: 60},
						{Uptime: 120},
						{Uptime: 0},
						{Uptime: 60},
						{Uptime: 0},
					},
				},
			}, nil
		},
	}

	expected := (3.0 / 5.0) * 100.0

	service := NewRouterService(mockRepo)
	result, err := service.Availability()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data := result["Router 1"]; data != float32(expected) {
		t.Errorf("expected %v, got %v", expected, data)
	}
}

func TestAvailability_LF(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 0},
						{Uptime: 60},
						{Uptime: 120},
						{Uptime: 0},
						{Uptime: 180},
						{Uptime: 0},
					},
				},
			}, nil
		},
	}

	expected := (4.0 / 5.0) * 100.0

	service := NewRouterService(mockRepo)
	result, err := service.Availability()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data := result["Router 1"]; data != float32(expected) {
		t.Errorf("expected %v, got %v", expected, data)
	}
}

func TestAvailability_OneData(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 0},
					},
				},
			}, nil
		},
	}

	expected := 0

	service := NewRouterService(mockRepo)
	result, err := service.Availability()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data := result["Router 1"]; data != float32(expected) {
		t.Errorf("expected %v, got %v", expected, data)
	}
}

// End of Availability Test

// Upload JSON Test
func TestProcessUploadedJSON_LoadError(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return nil, errors.New("failed to load json")
		},
	}

	service := NewRouterService(mockRepo)
	jsonFile := `
	[
		{
			"name": "Router Baru 1",
			"datas": [
			{ "uptime": 50 },
			{ "uptime": 60 },
			{ "uptime": 0 },
			{ "uptime": 120 }
			]
		}
	]
	`
	uploaded := []byte(jsonFile)

	error := service.ProcessUploadedJSON(uploaded)
	if error == nil {
		t.Error("Expected error, got nil")
	}
}

func TestProcessUploadedJSON_UnmarshallError(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 0},
						{Uptime: 60},
						{Uptime: 120},
					},
				},
			}, nil
		},
	}

	service := NewRouterService(mockRepo)
	jsonFile := `
	[
		{
			"name": "Router Baru 1",
			"datas": [
			{ "uptime": 50 },
			{ "uptime": 60 },
			{ "uptime": 0 },
			{ "uptime": 120
			]
		}
	]
	`
	uploaded := []byte(jsonFile)

	error := service.ProcessUploadedJSON(uploaded)
	if error == nil {
		t.Error("Expected error, got nil")
	}
}

func TestProcessUploadedJSON_Success(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return []models.Router{
				{
					Name: "Router 1",
					Datas: []models.Data{
						{Uptime: 0},
						{Uptime: 60},
						{Uptime: 120},
					},
				},
			}, nil
		},
		SaveFunc: func(r []models.Router) error {
			// Karena MarshalIndent perlu dua item di slice
			if len(r) != 2 {
				t.Errorf("Expected 2 routers, got %d", len(r))
			}
			return nil
		},
	}

	service := NewRouterService(mockRepo)
	jsonFile := `
	[
		{
			"name": "Router Baru 1",
			"datas": [
			{ "uptime": 50 },
			{ "uptime": 60 },
			{ "uptime": 0 },
			{ "uptime": 120}
			]
		}
	]
	`
	uploaded := []byte(jsonFile)

	error := service.ProcessUploadedJSON(uploaded)
	if error != nil {
		t.Error("Expected success")
	}
}
