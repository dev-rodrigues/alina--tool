package dto

import (
	"alina-tools/internal/domain"
)

type ServiceDTO struct {
	TotalPods int `json:"total_pods"`
}

func MapProcessDTOToService(dto ServiceDTO) domain.Service {
	return domain.Service{
		TotalPods: dto.TotalPods,
	}
}
