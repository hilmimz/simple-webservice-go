package handlers

import (
	"encoding/json"
	"net/http"
	"simple-webservice/internal/services"
)

/*
Sebenarnya kita bisa saja menuliskan bukan sebagai variabel pointer pada struct
Namun ini best practice terutama jika ada handler yang mengubah data asli
ataupun struct dengan ukuran besar
*/

type RouterHandler struct {
	Service *services.RouterService
}

func NewRouterHandler(s *services.RouterService) *RouterHandler {
	return &RouterHandler{Service: s}
}

func (h *RouterHandler) AvgUptime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := h.Service.AvgUptime()

	json.NewEncoder(w).Encode(result)
}

func (h *RouterHandler) Availability(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := h.Service.Availability()

	json.NewEncoder(w).Encode(result)
}
