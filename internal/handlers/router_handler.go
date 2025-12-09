package handlers

import (
	"encoding/json"
	"net/http"
	"simple-webservice/internal/services"
)

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
