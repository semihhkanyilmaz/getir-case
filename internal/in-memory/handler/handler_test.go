package inMemoryHandler

import (
	"bytes"
	inMemoryRepository "getir-case/internal/in-memory/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_ShouldBeReturnOK_GetRecordByKey(t *testing.T) {

	inMemRepo := inMemoryRepository.NewRepository()

	inMemRepo.Set("john", "doe")

	h := NewHandler(inMemRepo)

	r := httptest.NewRequest(http.MethodGet, "/api/in-memory?key=john", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Response code failed. Expected 200, actual %d", w.Code)
	}

	if strings.Trim(w.Body.String(), " ") != `{"value":"doe"}` {
		t.Errorf("Response body failed. Actual %s", w.Body.String())
	}
}

func Test_ShouldBeReturnNotFound_GetRecordByKey(t *testing.T) {
	inMemRepo := inMemoryRepository.NewRepository()

	h := NewHandler(inMemRepo)

	r := httptest.NewRequest(http.MethodGet, "/api/in-memory?key=john", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Response code failed. Expected 404, actual %d", w.Code)
	}

	if strings.Trim(w.Body.String(), " ") != `{"message":"john not found"}` {
		t.Errorf("Response body failed. Actual %s", w.Body.String())
	}
}

func Test_ShouldBeReturnBadRequest_GetRecordByKey(t *testing.T) {
	inMemRepo := inMemoryRepository.NewRepository()

	h := NewHandler(inMemRepo)

	r := httptest.NewRequest(http.MethodGet, "/api/in-memory", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Response code failed. Expected 400, actual %d", w.Code)
	}

	if strings.Trim(w.Body.String(), " ") != `{"message":"Param named key cannot be empty"}` {
		t.Errorf("Response body failed. Actual %s", w.Body.String())
	}
}

func Test_ShouldBeReturnCreated_CreateRecord(t *testing.T) {
	inMemRepo := inMemoryRepository.NewRepository()

	h := NewHandler(inMemRepo)

	payload := bytes.NewReader([]byte(`{"key":"john","value":"reader"}`))

	r := httptest.NewRequest(http.MethodPost, "/api/in-memory", payload)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Response code failed. Expected 201, actual %d", w.Code)
	}

	if strings.Trim(w.Body.String(), " ") != `` {
		t.Errorf("Response body failed. Actual %s", w.Body.String())
	}
}

//empty key
func Test_ShouldBeReturnBadRequest1_CreateRecord(t *testing.T) {
	inMemRepo := inMemoryRepository.NewRepository()

	h := NewHandler(inMemRepo)

	payload := bytes.NewReader([]byte(`{"key":"","value":"reader"}`))

	r := httptest.NewRequest(http.MethodPost, "/api/in-memory", payload)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Response code failed. Expected 400, actual %d", w.Code)
	}

	if strings.Trim(w.Body.String(), `{"message":"key cannot be empty"}`) != `` {
		t.Errorf("Response body failed. Actual %s", w.Body.String())
	}
}

//empty value
func Test_ShouldBeReturnBadRequest2_CreateRecord(t *testing.T) {
	inMemRepo := inMemoryRepository.NewRepository()

	h := NewHandler(inMemRepo)

	payload := bytes.NewReader([]byte(`{"key":"john","value":""}`))

	r := httptest.NewRequest(http.MethodPost, "/api/in-memory", payload)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Response code failed. Expected 400, actual %d", w.Code)
	}

	if strings.Trim(w.Body.String(), `{"message":"value cannot be empty"}`) != `` {
		t.Errorf("Response body failed. Actual %s", w.Body.String())
	}
}
