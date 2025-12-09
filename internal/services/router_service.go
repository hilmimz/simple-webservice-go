package services

import (
	"encoding/json"
	"simple-webservice/internal/helper"
	"simple-webservice/internal/models"
)

type RouterService struct {
	Routers []models.Router
}

/*
Sebenarnya kita bisa saja menuliskan bukan sebagai variabel pointer pada kembalian fungsi
Ini dilakukan karena router handler membutuhkan variabel pointer
Namun ini best practice terutama jika ada service yang mengubah data asli
ataupun struct dengan ukuran besar
*/
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

func (s *RouterService) Availability() map[string]float32 {
	data := make(map[string]float32)

	for _, router := range s.Routers {
		var availability int
		var unavailability int
		var isAnulir bool
		for _, data := range router.Datas {
			if data.Uptime > 0 {
				availability++
				if isAnulir {
					if data.Uptime > 60 {
						availability += unavailability
					}
					unavailability = 0
					isAnulir = false
				}
			} else {
				unavailability++
				isAnulir = true
			}
		}
		totalData := len(router.Datas)
		if isAnulir {
			totalData -= unavailability
		}
		data[router.Name] = float32(availability) / float32(totalData) * 100
	}
	return data
}
