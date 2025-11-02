package service

import (
	"testing"
)

func TestService_Convert(t *testing.T) {
	s := New()

	// Test Morse to text
	text, err := s.Convert(".--. .-. .. .-- . -") // ПРИВЕТ
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if text != "ПРИВЕТ" {
		t.Errorf("Expected 'ПРИВЕТ', got '%s'", text)
	}

	// Test text to Morse
	morse, err := s.Convert("ПРИВЕТ")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if morse != ".--. .-. .. .-- . -" {
		t.Errorf("Expected '.--. .-. .. .-- . -', got '%s'", morse)
	}
}
