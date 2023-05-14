package app

import (
	"regexp"
)

// The individual lex functions return an empty string if they failed to read a
// token.

var digitRegex = regexp.MustCompile("[0-9]")

func lexNumber(input string, i int) string {
	token := ""

	for j := i; j < len(input); j++ {
		charString := input[j : j+1]

		if digitRegex.MatchString(charString) {
			token += charString
		} else {
			return ""
		}
	}

	return token
}

func lexBoolean(input string, i int) string {
	if i+5 <= len(input) && input[i:i+5] == "false" {
		return "false"
	}

	if i+4 <= len(input) && input[i:i+4] == "true" {
		return "true"
	}

	return ""
}

// Does not currently handle escaped quotes
// Does not currently handle newlines
func lexString(input string, i int) string {
	j := i

	if input[j] != '"' {
		return ""
	}

	j++

	for ; j < len(input); j++ {
		if input[j] == '"' {
			return input[i : j+1]
		}
	}

	// We got to the end without finding a closing quote
	return ""
}

func Lex(input string) (tokens []string, ok bool) {
	tokens = []string{}
	i := 0

	check := func(token string) bool {
		if len(token) == 0 {
			return false
		}

		tokens = append(tokens, token)
		i += len(token)
		return true
	}

	for i < len(input) {
		if check(lexNumber(input, i)) {
			continue
		}

		if check(lexBoolean(input, i)) {
			continue
		}

		if check(lexString(input, i)) {
			continue
		}

		return tokens, false
	}

	return tokens, true
}
