package http

import (
	"alina-tools/app/pod/dto"
	"alina-tools/app/pod/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ProcessHandler struct {
	service *service.ProcessService
}

func NewProcessHandler(service *service.ProcessService) *ProcessHandler {
	return &ProcessHandler{service: service}
}

func (h *ProcessHandler) CreateService(context *gin.Context) {
	podId := context.Param("pod")
	var serviceDTO dto.ServiceDTO
	err := context.BindJSON(&serviceDTO)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return

	}

	service := dto.MapProcessDTOToService(serviceDTO)
	logrus.Info("Received request to create process: ", service)

	err, response := h.service.StartService(podId, service)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	context.JSON(http.StatusOK, response)
}

func (h *ProcessHandler) GetServices(context *gin.Context) {
	err, response := h.service.GetServices()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	context.JSON(http.StatusOK, response)
}

func (h *ProcessHandler) GetService(context *gin.Context) {
	podId := context.Param("pod")

	err, response := h.service.GetService(podId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	context.JSON(http.StatusOK, response)
}
