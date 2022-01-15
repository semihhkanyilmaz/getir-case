package router

import (
	inMemoryHandler "getir-case/internal/in-memory/handler"
	inMemoryRepository "getir-case/internal/in-memory/repository"
	recordHandler "getir-case/internal/record/handler"
	recordRepository "getir-case/internal/record/repository"
	"net/http"
)

func InitRoutes(inMemoryRepository inMemoryRepository.Repository, recordRepository recordRepository.Repository) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/in-memory", inMemoryHandler.NewHandler(inMemoryRepository))
	mux.Handle("/api/records", recordHandler.NewHandler(recordRepository))

	return mux
}
