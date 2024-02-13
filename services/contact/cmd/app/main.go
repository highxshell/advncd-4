package main

import (
	"architecture_go/pkg/store/postgres"
	delivery_http "architecture_go/services/contact/internal/delivery/http"
	postgres2 "architecture_go/services/contact/internal/repository/storage/postgres"
	"architecture_go/services/contact/internal/useCase/contact"
	"architecture_go/services/contact/internal/useCase/group"
	"log"
)

func main() {
	storage, err := postgres.New("localhost", "postgres", "Raider66", "postgres", 5432)
	if err != nil {
		log.Fatal("failed to init storage", err)
	}
	contactRepo := postgres2.NewContactRepository(storage)
	contactUC := contact.NewContactUseCase(contactRepo)
	groupUC := group.NewGroupUseCase(contactRepo, contactRepo)
	delivery := delivery_http.NewContactHTTP(contactUC, groupUC)
	delivery.Start(8080)
}
