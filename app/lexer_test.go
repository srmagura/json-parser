package app

import (
	"testing"
)

func assertFail(t *testing.T, json string) {
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

func TestInteger(t *testing.T) {
	assertTokensEqual(t, "1", []Token{{Number, "1"}})
	assertTokensEqual(t, "12", []Token{{Number, "12"}})
	assertTokensEqual(t, "123", []Token{{Number, "123"}})

	assertFail(t, "1x3")
}

func TestBoolean(t *testing.T) {
	assertTokensEqual(t, "false", []Token{{Boolean, "false"}})
	assertTokensEqual(t, "true", []Token{{Boolean, "true"}})
}

func TestString(t *testing.T) {
	assertTokensEqual(t, `"abc"`, []Token{{String, `"abc"`}})
}

func TestArray(t *testing.T) {
	assertTokensEqual(t, `[0, 1]`, []Token{
		{ArrayStart, "["},
		{Number, "0"},
		{Comma, ","},
		{Number, "1"},
		{ArrayEnd, "]"},
	})
}
