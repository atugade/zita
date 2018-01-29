package main

import (
	"strings"
)

func string_to_list(message string) []string {
	return strings.Split(message, " ")
}

func pop_list(message []string) []string {
	return append(message[:0], message[0+1:]...)
}
