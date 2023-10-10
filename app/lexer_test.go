package app

import (
	"testing"
)

func assertLexingFails(t *testing.T, json string) {
	_, ok := Lex(json)

	if ok {
		t.Fatal("Expected lexing to fail, but it succeeded.")
	}
}

func areTokensEqual(expected []Token, actual []Token) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			return false
		}
	}

	return true
}

func assertTokensEqual(t *testing.T, json string, expected []Token) {
	actual, ok := Lex(json)

	if !ok {
		t.Fatal("Lexing failed.")
	}

	if !areTokensEqual(expected, actual) {
		t.Error("Token arrays are not equal.")
		t.Error()
		t.Error("Expected:")
		t.Error(expected)
		t.Error()
		t.Error("Actual:")
		t.Error(actual)
		t.FailNow()
	}
}
