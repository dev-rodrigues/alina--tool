package dto

import (
	"alina-tools/internal/domain"
	"github.com/google/uuid"
)

type PodDTO struct {
	Name     string   `json:"name"`
	GitUrl   string   `json:"git_url"`
	Branch   string   `json:"branch"`
	Location string   `json:"location"`
	Commands []string `json:"commands"`
}

func MapPodDTOToPod(dto PodDTO) domain.Pod {
	return domain.Pod{
		ID:       uuid.New().String(),
		Name:     dto.Name,
		GitUrl:   dto.GitUrl,
		Branch:   dto.Branch,
		Location: dto.Location,
		Commands: dto.Commands,
		Started:  false,
		Pid:      0,
	}
}
