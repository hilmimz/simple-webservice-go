package main

import (
	"log"
	"net/http"
	"simple-webservice/internal/handlers"
	"simple-webservice/internal/middlewares"
	"simple-webservice/internal/services"
)

func main() {
	routerService, err := services.NewRouterService()

	if err != nil {
		panic(err)
	}

	routerHandler := handlers.NewRouterHandler(routerService)

	http.HandleFunc("/api/uptime/avg", middlewares.OnlyGET(routerHandler.AvgUptime))
	http.HandleFunc("/api/uptime/availability", middlewares.OnlyGET(routerHandler.Availability))

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
