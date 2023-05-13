package main

import (
	"regexp"
)

// The individual lex functions return an empty string if they failed to read a
// token.

func lexNumber(input string, i int) string {
	regex := regexp.MustCompile("[0-9]")

	token := ""

	for j := i; j < len(input); j++ {
		char := input[j : j+1]

		if regex.MatchString(char) {
			token += char
		} else {
			return ""
		}
	}

	return token
}

func Lex(input string) []string {
	tokens := []string{}

	i := 0
	token := lexNumber(input, i)

	if len(token) != 0 {
		tokens = append(tokens, token)
	}

	return tokens
}
