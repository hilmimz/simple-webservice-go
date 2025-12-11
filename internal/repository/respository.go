package repository

import (
	"encoding/json"
	"os"
	"simple-webservice/internal/models"
)

type RouterRepository interface {
	LoadRouters() ([]models.Router, error)
	SaveRouters([]models.Router) error
}

type JSONRouterRepository struct {
	FilePath string
}

func NewJSONRouterRepository(path string) *JSONRouterRepository {
	return &JSONRouterRepository{FilePath: path}
}

func (r *JSONRouterRepository) LoadRouters() ([]models.Router, error) {
	data, err := os.ReadFile(r.FilePath)

	if err != nil {
		return nil, err
	}

	var routers []models.Router
	if err := json.Unmarshal(data, &routers); err != nil {
		return nil, err
	}

	return routers, nil
}

func (r *JSONRouterRepository) SaveRouters(routers []models.Router) error {
	data, err := json.MarshalIndent(routers, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(r.FilePath, data, 0644)
}
