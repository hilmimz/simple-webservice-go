package services

import (
	"encoding/json"
	"simple-webservice/internal/helper"
	"simple-webservice/internal/models"
)

type RouterService struct {
	Routers []models.Router
}

func NewRouterService() (*RouterService, error) {
	data := helper.ReadFile("data/routers-uptime.json")

	var routers []models.Router
	if err := json.Unmarshal(data, &routers); err != nil {
		return nil, err
	}

	return &RouterService{Routers: routers}, nil
}

func (s *RouterService) AvgUptime() map[string]float32 {
	data := make(map[string]float32)

	for _, router := range s.Routers {
		var total int
		for _, data := range router.Datas {
			total += data.Uptime
		}
		avg := float32(total) / float32(len(router.Datas))
		data[router.Name] = avg
	}

	return data
}
