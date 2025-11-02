package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func New(logger *log.Logger, h *handlers.Handlers) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.RootHandler)
	mux.HandleFunc("/upload", h.UploadHandler)

	return &Server{
		logger: logger,
		httpServer: &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	s.logger.Println("Server is listening on", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
