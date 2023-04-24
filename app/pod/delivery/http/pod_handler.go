package http

import (
	"alina-tools/app/pod/dto"
	"alina-tools/app/pod/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PodHandler struct {
	service *service.PodService
}

func NewPodHandler(service *service.PodService) *PodHandler {
	return &PodHandler{service: service}
}

func (h *PodHandler) CreatePod(c *gin.Context) {
	var podDTO dto.PodDTO
	err := c.BindJSON(&podDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var pod = dto.MapPodDTOToPod(podDTO)
	logrus.Info("Received request to create pod: ", pod)

	err, d := h.service.CreatePod(pod)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d)
}

func (h *PodHandler) GetPods(context *gin.Context) {
	err, pods := h.service.GetPods()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, pods)
}

func (h *PodHandler) GetPod(context *gin.Context) {
	id := context.Param("id")
	err, pod := h.service.GetPod(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, pod)
}
