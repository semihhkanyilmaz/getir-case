package cmd

import (
	"context"
	inMemoryRepository "getir-case/internal/in-memory/repository"
	recordRepository "getir-case/internal/record/repository"
	router "getir-case/internal/router"
	"getir-case/pkg/helper"
	mongo "getir-case/pkg/mongo"
	"log"
	"net/http"
)

func Execute() {

	env := helper.ReadEnv()

	cfg := helper.ReadConfig(env)

	log.Printf("Http server started on %s", cfg.Port)
	log.Printf("Application running on : %s", env)

	mc := mongo.MewMongoClient(cfg.MongoSettings.ConnectionURI)

	inMemoryRepository := inMemoryRepository.NewRepository()
	recordRepository := recordRepository.NewRepository(mc.Database(cfg.MongoSettings.Database).Collection(cfg.MongoSettings.Collection))
	server := &http.Server{Addr: ":" + cfg.Port}
	server.Handler = router.InitRoutes(inMemoryRepository, recordRepository)

	if err := server.ListenAndServe(); err != nil {
		mc.Disconnect(context.Background())
		log.Fatal(err)
	}

}
