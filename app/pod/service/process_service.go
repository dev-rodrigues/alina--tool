package service

import (
	"alina-tools/app/pod/repository"
	"alina-tools/internal/domain"
	process2 "github.com/shirou/gopsutil/process"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type ProcessService struct {
	pod *repository.PodRepository
	cmd *repository.CmdMap
}

func NewProcessService(pod *repository.PodRepository, cmd *repository.CmdMap) *ProcessService {
	return &ProcessService{pod: pod, cmd: cmd}
}

func (s *ProcessService) StartService(podId string, service domain.Service) (error, *domain.Pod) {
	err, pod := s.pod.GetPod(podId)
	logrus.Info("Starting process for pod: with service:", podId, service)

	if err != nil {
		return err, nil
	}

	cmdMap := s.cmd.GetCmdMapInstance()
	cmdArgs := strings.Join(pod.Commands, " ")
	cmd := exec.Command("sh", "-c", cmdArgs)
	err = cmd.Start()

	pid := cmd.Process.Pid

	if err != nil {
		return err, nil
	} else {
		cmdMap.Map[podId] = cmd
	}

	pod.Started = true
	pod.Pid = pid

	err, d := s.pod.SavePod(*pod)
	if err != nil {
		return err, nil
	}

	return nil, d
}

func (s *ProcessService) GetServices() (error, []domain.Service) {
	services := s.cmd.GetCmdMapInstance().Map

	response := make([]domain.Service, len(services))

	for key, _ := range services {
		_, p := s.pod.GetPod(key)

		response = append(response, domain.Service{
			PodID:     p.ID,
			TotalPods: 1,
		})
	}

	return nil, response
}

func (s *ProcessService) GetService(id string) (error, *domain.ServiceDetails) {
	err, pod := s.pod.GetPod(id)
	if err != nil {
		logrus.Error(err)
		return err, nil
	}

	process, err := process2.NewProcess(int32(pod.Pid))
	if err != nil {
		logrus.Error(err)
		return err, nil
	}

	info, err := process.Status()
	if err != nil {
		logrus.Error(err)
		return err, nil
	}

	cpyPercentage, err := process.CPUPercent()
	if err != nil {
		logrus.Error(err)
		return err, nil
	}

	memoryPercentage, err := process.MemoryInfo()
	if err != nil {
		logrus.Error(err)
		return err, nil
	}

	createdTime, err := process.CreateTime()
	if err != nil {
		logrus.Error(err)
		return err, nil
	}

	details := domain.ServiceDetails{
		PodID:         id,
		TotalPods:     1,
		Name:          pod.Name,
		State:         mapStatusToMeaning(info),
		CPUPercent:    cpyPercentage,
		MemoryPercent: memoryPercentage.RSS,
		CreatedAt:     createdTime,
	}

	return nil, &details
}

func mapStatusToMeaning(status string) string {
	meaningMap := map[string]string{
		"R": "running",
		"S": "sleeping",
		"D": "disk sleep",
		"Z": "zombie",
		"T": "stopped",
		"t": "tracing",
		"X": "dead",
		"x": "dead and wakeable",
	}

	if meaning, ok := meaningMap[status]; ok {
		return meaning
	}

	return "Error"
}
