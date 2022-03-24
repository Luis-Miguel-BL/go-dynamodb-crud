package services

import (
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/repositories"
)

type HealthService struct {
	repository repositories.HealthRepository
}

type HealthInterface interface {
	Health() bool
}

func NewHealthService(repository repositories.HealthRepository) HealthInterface {
	return &HealthService{repository: repository}
}

func (h *HealthService) Health() bool {
	return h.repository.Health()
}
