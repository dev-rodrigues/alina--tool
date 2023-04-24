package main

import (
	"alina-tools/app/pod/delivery/http"
	"alina-tools/app/pod/repository"
	"alina-tools/app/pod/service"
	"alina-tools/internal/database"
	"alina-tools/internal/log"
	"alina-tools/internal/server"
)

func main() {

	log.ConfigureLogger()

	redisClient := database.NewRedisClient()
	defer redisClient.Close()

	podRepository := repository.NewPodRepository(redisClient)

	podService := service.NewPodService(podRepository)
	podHandler := http.NewPodHandler(podService)

	pm2Handler := http.NewPm2Handler()

	router := server.NewRouter()

	router.POST("/pod", podHandler.CreatePod)
	router.GET("/pod", podHandler.GetPods)
	router.GET("/pod/:id", podHandler.GetPod)

	router.GET("/pm2/list", pm2Handler.GetPm2List)

	router.Run(":8080")
}
