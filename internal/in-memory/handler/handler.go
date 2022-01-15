package inMemoryHandler

import (
	inMemoryType "getir-case/internal/in-memory/model"
	inMemoryRepository "getir-case/internal/in-memory/repository"
	errors "getir-case/pkg/error"
	httpContext "getir-case/pkg/http-context"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	inMemoryRepository inMemoryRepository.Repository
}

func NewHandler(inMemoryRepository inMemoryRepository.Repository) *Handler {
	return &Handler{
		inMemoryRepository: inMemoryRepository,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := httpContext.NewContext(w, r)
	defer c.Recover()

	switch r.Method {
	case http.MethodPost:
		h.createRecord(c)
	case http.MethodGet:
		h.getRecordByKey(c)
	default:
		c.Json(http.StatusMethodNotAllowed, errors.NewError(http.StatusText(http.StatusMethodNotAllowed)))
	}

}

func (h *Handler) createRecord(c httpContext.Context) {

	record := new(inMemoryType.PostInMemoryRequest)

	if err := c.Bind(record); err != nil {
		go log.Println(err)
		c.Json(http.StatusBadRequest, errors.NewError("Invalid payload"))
		return
	}

	if strings.TrimSpace(record.Key) == "" {
		c.Json(http.StatusBadRequest, errors.NewError("key cannot be empty"))
	}
	if strings.TrimSpace(record.Value) == "" {
		c.Json(http.StatusBadRequest, errors.NewError("value cannot be empty"))
	}

	h.inMemoryRepository.Set(record.Key, record.Value)

	c.NoContent(http.StatusCreated)
}

func (h *Handler) getRecordByKey(c httpContext.Context) {

	key := c.GetQueryParam("key")

	if strings.TrimSpace(key) == "" {
		c.Json(http.StatusBadRequest, errors.NewError("Param named key cannot be empty"))
		return
	}

	value, err := h.inMemoryRepository.Get(key)
	if err != nil {
		go log.Println(err)
		c.Json(http.StatusNotFound, errors.NewError(err.Error()))
		return
	}

	c.Json(http.StatusOK, inMemoryType.GetInMemoryResponse{
		Value: value,
	})
}
