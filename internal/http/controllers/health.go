package controllers

import (
	"errors"
	"net/http"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/services"
	HttpStatus "github.com/Luis-Miguel-BL/go-dynamodb-crud/utils/http"
)

type HealthController struct {
	Service services.HealthInterface
}

func NewHealthController(service services.HealthInterface) *HealthController {
	return &HealthController{
		Service: service,
	}
}

func (h *HealthController) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Service.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}
