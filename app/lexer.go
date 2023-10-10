package app

import (
	"regexp"
)

type TokenType byte

const (
	TokenTypeNumber TokenType = iota
	TokenTypeBoolean
	TokenTypeString
	TokenTypeArrayStart
	TokenTypeArrayEnd
	TokenTypeComma
)

type Token struct {
	Type  TokenType
	Value string
}

// The individual lex functions return an empty string if they failed to read a
// token.

var digitRegex = regexp.MustCompile("[0-9]")

func lexNumber(input string, i int) *Token {
	tokenValue := ""

	// This is going to need to be reworked to support floats
	for j := i; j < len(input); j++ {
		charString := input[j : j+1]

		if digitRegex.MatchString(charString) {
			tokenValue += charString
		} else {
			break
		}
	}

	if len(tokenValue) == 0 {
		return nil
	}

	return &Token{TokenTypeNumber, tokenValue}
}

func lexBoolean(input string, i int) *Token {
	if i+5 <= len(input) && input[i:i+5] == "false" {
		return &Token{TokenTypeBoolean, "false"}
	}

	if i+4 <= len(input) && input[i:i+4] == "true" {
		return &Token{TokenTypeBoolean, "true"}
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
			return &Token{TokenTypeString, input[i : j+1]}
		}
	}

	// We got to the end without finding a closing quote
	return nil
}

func lexDelimiter(input string, i int) *Token {
	if input[i] == ',' {
		return &Token{TokenTypeComma, ","}
	}

	if input[i] == '[' {
		return &Token{TokenTypeArrayStart, "["}
	}

	if input[i] == ']' {
		return &Token{TokenTypeArrayEnd, "]"}
	}

	// TODO object start/end

	return nil
}

func Lex(input string) (tokens []Token, ok bool) {
	tokens = []Token{}
	i := 0

	check := func(token *Token) bool {
		if token == nil {
			return false
		}

		tokens = append(tokens, *token)
		i += len(token.Value)
		return true
	}

	for i < len(input) {
		nextChar := input[i]

		// Spaces and tabs are not considered tokens
		if nextChar == ' ' || nextChar == '\t' {
			i++
			continue
		}

		if check(lexNumber(input, i)) {
			continue
		}

		if check(lexBoolean(input, i)) {
			continue
		}

		if check(lexString(input, i)) {
			continue
		}

		if check(lexDelimiter(input, i)) {
			continue
		}

		return []Token{}, false
	}

	return tokens, true
}
