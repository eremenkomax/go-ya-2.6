package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func TestHandlers_RootHandler(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}
	err = os.Chdir("../..")
	if err != nil {
		t.Fatalf("failed to change dir to root: %v", err)
	}
	defer os.Chdir(wd)

	s := service.New()
	h := New(s)

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	h.RootHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected, err := os.ReadFile("index.html")
	if err != nil {
		t.Fatalf("could not read index.html: %v", err)
	}
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}

func TestHandlers_UploadHandler(t *testing.T) {
	s := service.New()
	h := New(s)

	// Create a new file upload request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("myFile", "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.WriteString(part, "ПРИВЕТ")
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()

	h.UploadHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := ".--. .-. .. .-- . -"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHandlers_RootHandler_MethodNotAllowed(t *testing.T) {
	s := service.New()
	h := New(s)

	req := httptest.NewRequest("POST", "/", nil)
	rr := httptest.NewRecorder()

	h.RootHandler(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestHandlers_UploadHandler_MethodNotAllowed(t *testing.T) {
	s := service.New()
	h := New(s)

	req := httptest.NewRequest("GET", "/upload", nil)
	rr := httptest.NewRecorder()

	h.UploadHandler(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestHandlers_UploadHandler_Errors(t *testing.T) {
	s := service.New()
	h := New(s)

	// Test case 1: Invalid form
	req := httptest.NewRequest("POST", "/upload", bytes.NewReader([]byte("invalid")))
	req.Header.Set("Content-Type", "multipart/form-data; boundary=...")
	rr := httptest.NewRecorder()
	h.UploadHandler(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	// Test case 2: No file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()
	req = httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr = httptest.NewRecorder()
	h.UploadHandler(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
