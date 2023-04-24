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
	cmdRepository := repository.NewCmdMap()

	podService := service.NewPodService(podRepository)
	podHandler := http.NewPodHandler(podService)

	processService := service.NewProcessService(podRepository, cmdRepository)
	procHandler := http.NewProcessHandler(processService)

	router := server.NewRouter()

	router.POST("/pod", podHandler.CreatePod)
	router.GET("/pod", podHandler.GetPods)
	router.GET("/pod/:id", podHandler.GetPod)

	router.GET("/services", procHandler.GetServices)
	router.POST("/services/:pod", procHandler.CreateService)
	router.GET("/services/:pod", procHandler.GetService)

	router.Run(":8080")
}
