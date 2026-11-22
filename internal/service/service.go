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

	isMorse := !strings.ContainsFunc(s, func(r rune) bool {
		return r != '.' && r != '-' && r != ' ' && r != '\n' && r != '\r'
	})
	if isMorse {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}
