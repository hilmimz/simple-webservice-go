package handlers

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"simple-webservice/internal/models"
	"simple-webservice/internal/services"
	"simple-webservice/internal/test"
	"strings"
	"testing"
)

// type errorFile struct{}

// func (e *errorFile) Read(p []byte) (int, error) {
// 	return 0, errors.New("read error")
// }

// func (e *errorFile) ReadAt(p []byte, off int64) (int, error) {
// 	return 0, errors.New("read error")
// }

// func (e *errorFile) Seek(offset int64, whence int) (int64, error) {
// 	return 0, nil
// }

// func (e *errorFile) Close() error {
// 	return nil
// }

// type errorFileHeader struct {
// 	multipart.FileHeader
// }

// func (h *errorFileHeader) Open() (multipart.File, error) {
// 	return &errorFile{}, nil
// }

func TestUpload_Success(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("routerData", "routers.json")
	part.Write([]byte(`
	[
		{
			"name": "Router 1",
			"datas": [
			{ "uptime": 0 },
			{ "uptime": 60 },
			{ "uptime": 120 }
			]
		}
	]
	`))
	writer.Close()

	req := httptest.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// mock := &test.MockRouterService{}
	// service := services.NewRouterService(mock)
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return nil, nil
		},
		SaveFunc: func([]models.Router) error {
			return nil
		},
	}
	service := services.NewRouterService(mockRepo)
	handler := NewRouterHandler(service)

	rr := httptest.NewRecorder()

	handler.UploadHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "uploaded and merged successful") {
		t.Fatalf("unexpected body: %s", rr.Body.String())
	}
}

func TestUpload_NoFile(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req := httptest.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	mock := &test.MockRouterService{}
	service := services.NewRouterService(mock)
	handler := NewRouterHandler(service)

	rr := httptest.NewRecorder()

	handler.UploadHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestUpload_MergeJSONError(t *testing.T) {
	mockRepo := &test.MockRouterService{
		LoadFunc: func() ([]models.Router, error) {
			return nil, nil
		},
		SaveFunc: func([]models.Router) error {
			return errors.New("invalid json")
		},
	}
	service := services.NewRouterService(mockRepo)
	handler := NewRouterHandler(service)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("routerData", "routers.json")
	part.Write([]byte(`[{ "name":"Router X"}]`))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()

	handler.UploadHandler(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rr.Code)
	}
}

// func TestUpload_ReadError(t *testing.T) {
// 	mockRepo := &test.MockRouterService{
// 		LoadFunc: func() ([]models.Router, error) {
// 			return nil, nil
// 		},
// 		SaveFunc: func([]models.Router) error {
// 			return nil
// 		},
// 	}
// 	service := services.NewRouterService(mockRepo)
// 	handler := NewRouterHandler(service)

// 	// buat mock FileHeader
// 	mockFH := &errorFileHeader{
// 		FileHeader: multipart.FileHeader{
// 			Filename: "mock.json",
// 		},
// 	}

// 	// buat request dan inject MultipartForm manual
// 	req := httptest.NewRequest("POST", "/upload", nil)
// 	req.MultipartForm = &multipart.Form{
// 		File: map[string][]*multipart.FileHeader{
// 			"routerData": {&mockFH.FileHeader},
// 		},
// 	}

// 	rr := httptest.NewRecorder()

// 	handler.UploadHandler(rr, req)

// 	// assert
// 	if rr.Code != http.StatusInternalServerError {
// 		t.Fatalf("expected 400, got %d", rr.Code)
// 	}

// 	if !strings.Contains(rr.Body.String(), "failed to read content") {
// 		t.Fatalf("unexpected body: %s", rr.Body.String())
// 	}
// }
