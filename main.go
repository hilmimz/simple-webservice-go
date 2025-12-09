package main

import (
	"log"
	"net/http"
	"simple-webservice/internal/handlers"
	"simple-webservice/internal/services"
)

func main() {
	routerService, err := services.NewRouterService()

	if err != nil {
		panic(err)
	}

	routerHandler := handlers.NewRouterHandler(routerService)

	http.HandleFunc("/uptime/avg", routerHandler.AvgUptime)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
