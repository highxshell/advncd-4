package main

import (
	"architecture_go/pkg/store/postgres"
	delivery_http "architecture_go/services/contact/internal/delivery/http"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/useCase"
	"log"
)

func main() {
	storage, err := postgres.New("localhost", "postgres", "Raider66", "postgres", 5432)
	if err != nil {
		log.Fatal("failed to init storage", err)
	}
	contactRepo := repository.NewContactRepository(storage)
	contactUC := useCase.NewContactUseCase(contactRepo)
	delivery := delivery_http.NewContactHTTP(contactUC)
	delivery.Start(8080)
}
