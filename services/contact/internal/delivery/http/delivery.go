package http

import (
	"architecture_go/services/contact/internal/useCase"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type ContactHTTPDelivery struct {
	contactUseCase useCase.ContactUseCase
	groupUseCase   useCase.GroupUseCase
}

func NewContactHTTP(contactUseCase useCase.ContactUseCase, groupUseCase useCase.GroupUseCase) *ContactHTTPDelivery {
	return &ContactHTTPDelivery{contactUseCase: contactUseCase, groupUseCase: groupUseCase}
}

func (d *ContactHTTPDelivery) Start(port int) {
	const op = "delivery.http.Start"
	addr := fmt.Sprintf(":%d", port)
	go func() {
		log.Println("Starting server on port", port)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("%s: %w", op, err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down gracefully...")
}
