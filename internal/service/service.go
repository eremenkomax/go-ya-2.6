package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Service does something good.
type Service struct{}

// New returns a new Service.
func New() *Service {
	return &Service{}
}

// Convert auto-detects if the input is morse and converts it to the other.
func (s *Service) Convert(input string) (string, error) {
	if isMorse(input) {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}

func isMorse(s string) bool {
	s = strings.TrimSpace(s)
	if !strings.ContainsAny(s, ".-") {
		return false
	}
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			return false
		}
	}
	return true
}