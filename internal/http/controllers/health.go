package controllers

import (
	"errors"
	"net/http"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/repositories"
	HttpStatus "github.com/Luis-Miguel-BL/go-dynamodb-crud/utils/http"
)

type HealthController struct {
	Repository repositories.HealthRepository
}

func NewHealthController(repository repositories.HealthRepository) *HealthController {
	return &HealthController{
		Repository: repository,
	}
}

func (h *HealthController) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}
