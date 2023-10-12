package main

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

func TestLexNumber(t *testing.T) {
	assertTokensEqual(t, "1", []Token{{TokenTypeNumber, "1"}})
	assertTokensEqual(t, "12", []Token{{TokenTypeNumber, "12"}})
	assertTokensEqual(t, "123", []Token{{TokenTypeNumber, "123"}})

	assertTokensEqual(t, "123.", []Token{{TokenTypeNumber, "123."}})
	assertTokensEqual(t, "123.4", []Token{{TokenTypeNumber, "123.4"}})

	// These strings are not valid JSON, but they are sequences of valid tokens.
	// The failure will occur during parsing.
	assertTokensEqual(t, "1..3", []Token{{TokenTypeNumber, "1."}, {TokenTypeNumber, ".3"}})
	assertTokensEqual(t, "1.2.3", []Token{{TokenTypeNumber, "1.2"}, {TokenTypeNumber, ".3"}})

	assertLexingFails(t, "1x3")
}

func TestLexBoolean(t *testing.T) {
	assertTokensEqual(t, "false", []Token{{TokenTypeBoolean, "false"}})
	assertTokensEqual(t, "true", []Token{{TokenTypeBoolean, "true"}})
}

func TestLexString(t *testing.T) {
	assertTokensEqual(t, `"abc"`, []Token{{TokenTypeString, `"abc"`}})
}

func TestLexArray(t *testing.T) {
	assertTokensEqual(t, `[0, 1]`, []Token{
		{TokenTypeArrayStart, "["},
		{TokenTypeNumber, "0"},
		{TokenTypeComma, ","},
		{TokenTypeNumber, "1"},
		{TokenTypeArrayEnd, "]"},
	})
}
