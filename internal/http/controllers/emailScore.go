package controllers

import (
	"net/http"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/http/viewmodel"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/repositories"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/services"
	HttpStatus "github.com/Luis-Miguel-BL/go-dynamodb-crud/utils/http"
)

type EmailScoreController struct {
	Service services.EmailScoreInterface
}

func NewEmailScoreController(repository repositories.EmailScoreRepository) *EmailScoreController {
	return &EmailScoreController{
		Service: services.NewEmailScoreService(repository),
	}
}

func (h *EmailScoreController) FindByEmails(w http.ResponseWriter, r *http.Request) {

	emails, err := viewmodel.ValidateFindByEmails(r)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}
	emailScores, err := h.Service.ListByEmails(emails)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, emailScores)
}

func (h *EmailScoreController) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	response, err := h.Service.GetByEmail(email)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *EmailScoreController) ConsolidateScore(w http.ResponseWriter, r *http.Request) {
	consolidateScoreBody, err := viewmodel.ValidateConsolidateScore(r)

	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}
	err = h.Service.ConsolidateScore(consolidateScoreBody)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)
}
