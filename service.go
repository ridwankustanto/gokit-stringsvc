package main

import (
	"errors"
	"strings"
)

// StringService contain methods that take care of strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

// Business logic

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("String is empty")
