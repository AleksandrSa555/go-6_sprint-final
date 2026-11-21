package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("пустая строка")
	}

	if strings.Trim(s, ".- \n\r") == "" {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}
