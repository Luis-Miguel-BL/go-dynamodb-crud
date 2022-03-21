package health

import (
	"errors"
	"net/http"

	"github.com/akhil/dynamodb-go-crud-yt/internal/controllers"
	"github.com/akhil/dynamodb-go-crud-yt/internal/repository/adapter"
	HttpStatus "github.com/akhil/dynamodb-go-crud-yt/utils/http"
)

type Handler struct {
	controllers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) controllers.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
