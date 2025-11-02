package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	appService := service.New()
	appHandlers := handlers.New(appService)
	srv := server.New(logger, appHandlers)

	if err := srv.Run(); err != nil {
		logger.Fatalf("server stopped with error: %s", err)
	}
}