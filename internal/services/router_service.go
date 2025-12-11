package services

import (
	"encoding/json"
	"math"
	"simple-webservice/internal/models"
	"simple-webservice/internal/repository"
)

type RouterService struct {
	Repo repository.RouterRepository
}

/*
Sebenarnya kita bisa saja menuliskan bukan sebagai variabel pointer pada kembalian fungsi
Ini dilakukan karena router handler membutuhkan variabel pointer
Namun ini best practice terutama jika ada service yang mengubah data asli
ataupun struct dengan ukuran besar
*/
func NewRouterService(repo repository.RouterRepository) *RouterService {
	return &RouterService{Repo: repo}
}

func (s *RouterService) AvgUptime() (map[string]float32, error) {
	data := make(map[string]float32)
	routers, err := s.Repo.LoadRouters()
	if err != nil {
		return nil, err
	}

	for _, router := range routers {
		var total int
		for _, data := range router.Datas {
			total += data.Uptime
		}
		avg := float32(total) / float32(len(router.Datas))
		data[router.Name] = avg
	}

	return data, nil
}

func (s *RouterService) Availability() (map[string]float32, error) {
	data := make(map[string]float32)
	routers, err := s.Repo.LoadRouters()
	if err != nil {
		return nil, err
	}

	for _, router := range routers {
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
		if unavailability == totalData {
			data[router.Name] = float32(0)
		} else {
			if isAnulir {
				totalData -= unavailability
			}
			result := float32(availability) / float32(totalData) * 100
			data[router.Name] = float32(math.Round(float64(result)))
		}
	}
	return data, nil
}

func (s *RouterService) ProcessUploadedJSON(uploadedJSON []byte) error {
	existing, err := s.Repo.LoadRouters()
	if err != nil {
		return err
	}

	var uploaded []models.Router
	if err := json.Unmarshal(uploadedJSON, &uploaded); err != nil {
		return err
	}

	merged := append(existing, uploaded...)

	return s.Repo.SaveRouters(merged)
}
