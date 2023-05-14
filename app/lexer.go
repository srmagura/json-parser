package app

import (
	"regexp"
)

type TokenType byte

const (
	Number TokenType = iota
	Boolean
	String
)

type Token struct {
	Type  TokenType
	Value string
}

// The individual lex functions return an empty string if they failed to read a
// token.

var digitRegex = regexp.MustCompile("[0-9]")

func lexNumber(input string, i int) *Token {
	token := ""

	// This is going to need to be reworked to support floats
	for j := i; j < len(input); j++ {
		charString := input[j : j+1]

		if digitRegex.MatchString(charString) {
			token += charString
		} else {
			return nil
		}
	}

	return &Token{Number, token}
}

func lexBoolean(input string, i int) *Token {
	if i+5 <= len(input) && input[i:i+5] == "false" {
		return &Token{Boolean, "false"}
	}

	if i+4 <= len(input) && input[i:i+4] == "true" {
		return &Token{Boolean, "true"}
	}

	return nil
}

// Does not currently handle escaped quotes
// Does not currently handle newlines
func lexString(input string, i int) *Token {
	j := i

	if input[j] != '"' {
		return nil
	}

	j++

	for ; j < len(input); j++ {
		if input[j] == '"' {
			return &Token{String, input[i : j+1]}
		}
	}

	// We got to the end without finding a closing quote
	return nil
}

func Lex(input string) (tokens []string, ok bool) {
	tokens = []string{}
	i := 0

	check := func(token *Token) bool {
		if token == nil {
			return false
		}

		tokens = append(tokens, token.Value)
		i += len(token.Value)
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
