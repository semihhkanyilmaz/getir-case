package recordHandler

import (
	recordType "getir-case/internal/record/model"
	recordRepository "getir-case/internal/record/repository"
	errors "getir-case/pkg/error"
	httpContext "getir-case/pkg/http-context"
	"net/http"
)

type handler struct {
	recordRepository recordRepository.Repository
}

func NewHandler(recordRepository recordRepository.Repository) *handler {
	return &handler{
		recordRepository: recordRepository,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := httpContext.NewContext(w, r)
	defer c.Recover()

	switch r.Method {
	case http.MethodPost:
		h.getRecords(c)
	default:
		c.Json(http.StatusMethodNotAllowed, errors.NewError(http.StatusText(http.StatusMethodNotAllowed)))
	}

}

func (h *handler) getRecords(c httpContext.Context) {

	req := new(recordType.GetRecordsRequest)
	res := new(recordType.GetRecordsResponse)

	res.Records = make([]recordType.Record, 0)

	if err := c.Bind(req); err != nil {
		res.Msg = "Invalid payload"
		c.Json(http.StatusBadRequest, res)
		return
	}

	dbReq, err := req.ToGetRecordsDBModel()
	if err != nil {
		res.Msg = err.Error()
		c.Json(http.StatusBadRequest, res)
		return
	}

	data := h.recordRepository.Aggregate(dbReq)

	c.Json(http.StatusOK, data)
}
