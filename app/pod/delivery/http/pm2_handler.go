package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type Pm2Handler struct {
}

func NewPm2Handler() *Pm2Handler {
	return &Pm2Handler{}
}

func (h *Pm2Handler) GetPm2List(context *gin.Context) {

	cmd := exec.Command(
		"-c",
		"cd ~/Documents/workspace/faculdade/fakewaze/build/libs",
		"pm2",
		"list",
	)
	output, err := cmd.Output()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, string(output))
}
