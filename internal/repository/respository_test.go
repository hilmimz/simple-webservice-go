package repository

import (
	"errors"
	"os"
	"path/filepath"
	"simple-webservice/internal/models"
	"simple-webservice/internal/test"
	"testing"
)

func TestLoadRouters_Success(t *testing.T) {
	tempDir := t.TempDir()

	dummyJSON := `[
		{
			"name": "Router 1",
			"datas": [
			{ "uptime": 0 },
			{ "uptime": 60 },
			{ "uptime": 120 }
			]
		},
		{
			"name": "Router 2",
			"datas": [
			{ "uptime": 0 },
			{ "uptime": 0 },
			{ "uptime": 60 }
			]
		}
	]`

	filePath := filepath.Join(tempDir, "routers.json")
	if err := os.WriteFile(filePath, []byte(dummyJSON), 0644); err != nil {
		t.Fatal(err)
	}

	repo := NewJSONRouterRepository(filePath)
	routers, err := repo.LoadRouters()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(routers) != 2 {
		t.Errorf("expected 2 routers, got %d", len(routers))
	}

	if routers[0].Name != "Router 1" {
		t.Errorf("expected router name Router 1, got %s", routers[0].Name)
	}

}

func TestLoadRouters_NoFile(t *testing.T) {
	repo := NewJSONRouterRepository("/dummy")

	_, err := repo.LoadRouters()
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

func TestLoadRouters_UnmarshalError(t *testing.T) {
	tempDir := t.TempDir()

	dummyJSON := `[
		{
			"nam": "Router 1",
			"datas": [
			{ "uptime": 0 },
			{ "uptime": 60 },
			{ "uptime": 120 
			]
		}
	]`

	filePath := filepath.Join(tempDir, "routers.json")
	if err := os.WriteFile(filePath, []byte(dummyJSON), 0644); err != nil {
		t.Fatal(err)
	}

	mock := NewJSONRouterRepository(filePath)
	_, err := mock.LoadRouters()

	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

func TestSaveRouters_Success(t *testing.T) {
	tempDir := t.TempDir()

	filepath := filepath.Join(tempDir, "routers.json")

	mock := NewJSONRouterRepository(filepath)
	var routers []models.Router

	if err := mock.SaveRouters(routers); err != nil {
		t.Errorf("expected succes, got %v", err)
	}
}

func TestSaveRouters_MarshallIndentError(t *testing.T) {
	mockRepo := &test.MockRouterService{
		SaveFunc: func([]models.Router) error {
			return errors.New("marshal indent error")
		},
	}

	var routers []models.Router
	if err := mockRepo.SaveFunc(routers); err == nil {
		t.Errorf("expected error")
	}
}
