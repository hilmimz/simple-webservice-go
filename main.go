package main

import (
	"log"
	"net/http"
	"simple-webservice/internal/handlers"
	"simple-webservice/internal/middlewares"
	"simple-webservice/internal/repository"
	"simple-webservice/internal/services"
)

func main() {
	// routerService, err := services.NewRouterService()

	// if err != nil {
	// 	panic(err)
	// }

	filePath := "data/routers-uptime.json"

	//Buat Repo
	repo := repository.NewJSONRouterRepository(filePath)
	// Buat service
	routerService := services.NewRouterService(repo)

	// Buat handler
	routerHandler := handlers.NewRouterHandler(routerService)

	// routerHandler := handlers.RouterHandler{}

	http.HandleFunc("/api/uptime/avg", middlewares.OnlyGET(routerHandler.AvgUptime))
	http.HandleFunc("/api/uptime/availability", middlewares.OnlyGET(routerHandler.Availability))
	http.HandleFunc("/api/upload", routerHandler.UploadHandler)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
