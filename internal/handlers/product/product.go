package product

import (
	"net/http"

	"github.com/danish45007/go-dynamodb/internal/handlers"
	"github.com/danish45007/go-dynamodb/internal/repository/adapter"
	Http "github.com/danish45007/go-dynamodb/utils/http"
)

type Handler struct {
	handlers.Interface
	Contrtoller product.Interface
	Rules       rules.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}
func (h *Handler) getBodyAndValidate(request *http.Request) {}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetALL(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		Http.StatusBadRequest(w, r, err)
	}
	ID, err := h.Contrtoller.Create(productBody)
	if err != nil {
		Http.StatusInternalServerError(w, r, err)
		return
	}
	Http.StatusOk(w, r, map[string]interface{}{"id": ID.String()})
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {

}
