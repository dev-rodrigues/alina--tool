package service

import (
	repository "alina-tools/app/pod/repository"
	"alina-tools/internal/domain"
)

type PodService struct {
	repo *repository.PodRepository
}

func NewPodService(repo *repository.PodRepository) *PodService {
	return &PodService{repo: repo}
}

func (s *PodService) CreatePod(pod domain.Pod) (error, *domain.Pod) {
	return s.repo.SavePod(pod)
}

func (s *PodService) GetPods() (error, []domain.Pod) {
	return s.repo.GetPods()
}

func (s *PodService) GetPod(id string) (interface{}, interface{}) {
	return s.repo.GetPod(id)
}
