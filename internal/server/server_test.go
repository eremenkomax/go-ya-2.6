package server

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func TestNew(t *testing.T) {
	logger := log.New(os.Stdout, "", 0)
	s := service.New()
	h := handlers.New(s)
	srv := New(logger, h)

	if srv.httpServer.Addr != ":8080" {
		t.Errorf("expected address to be :8080, got %s", srv.httpServer.Addr)
	}
}

func TestServer_Run(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}
	err = os.Chdir("../..")
	if err != nil {
		t.Fatalf("failed to change dir to root: %v", err)
	}
	defer os.Chdir(wd)

	logger := log.New(os.Stdout, "", 0)
	s := service.New()
	h := handlers.New(s)
	srv := New(logger, h)

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			t.Errorf("server failed to run: %v", err)
		}
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Try to make a request to the server
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Errorf("failed to make request to server: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %s", resp.Status)
	}

	// Stop the server
	if err := srv.httpServer.Close(); err != nil {
		t.Errorf("failed to stop server: %v", err)
	}
}
