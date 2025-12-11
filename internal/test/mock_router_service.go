package test

import "simple-webservice/internal/models"

type MockRouterService struct {
	LoadFunc func() ([]models.Router, error)
	SaveFunc func([]models.Router) error
}

func (m *MockRouterService) LoadRouters() ([]models.Router, error) {
	return m.LoadFunc()
}

func (m *MockRouterService) SaveRouters(routers []models.Router) error {
	return m.SaveFunc(routers)
}
