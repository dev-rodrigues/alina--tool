package repository

import (
	"alina-tools/internal/domain"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type PodRepository struct {
	client *redis.Client
}

func NewPodRepository(client *redis.Client) *PodRepository {
	return &PodRepository{
		client: client,
	}
}

func (r PodRepository) SavePod(pod domain.Pod) (error, *domain.Pod) {
	ctx := context.Background()

	jsonPod, err := json.Marshal(pod)
	if err != nil {
		return err, nil
	}

	err = r.client.Set(ctx, pod.ID, jsonPod, 0).Err()
	if err != nil {
		return err, nil
	}

	return nil, &pod
}

func (r PodRepository) GetPods() (error, []domain.Pod) {
	ctx := context.Background()

	var pods []domain.Pod

	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return err, nil
	}

	for _, key := range keys {
		val, err := r.client.Get(ctx, key).Result()
		if err != nil {
			return err, nil
		}

		var pod domain.Pod
		err = json.Unmarshal([]byte(val), &pod)
		if err != nil {
			return err, nil
		}

		pods = append(pods, pod)
	}

	return nil, pods
}

func (r PodRepository) GetPod(id string) (error, *domain.Pod) {
	ctx := context.Background()

	val, err := r.client.Get(ctx, id).Result()
	if err != nil {
		return err, nil
	}

	var pod domain.Pod
	err = json.Unmarshal([]byte(val), &pod)
	if err != nil {
		return err, nil
	}

	return nil, &pod
}
