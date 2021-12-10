package health

import (
	"errors"
	"net/http"

	"github.com/danish45007/go-dynamodb/internal/handlers"
	"github.com/danish45007/go-dynamodb/internal/repository/adapter"
	Http "github.com/danish45007/go-dynamodb/utils/http"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	health := h.Repository.DbHealth()
	if !health {
		Http.StatusInternalServerError(w, r, errors.New("dynamodb not alive"))
		return
	}
	Http.StatusOk(w, r, "dynamodb live")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	Http.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	Http.StatusMethodNotAllowed(w, r)

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	Http.StatusMethodNotAllowed(w, r)

}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	Http.StatusMethodNotAllowed(w, r)

}
