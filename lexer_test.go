package main

import (
	"testing"
)

func AreTokensEqual(expected []string, actual []string) bool {
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

func AssertTokensEqual(t *testing.T, expected []string, actual []string) {
	if !AreTokensEqual(expected, actual) {
		t.Error("Token arrays are not equal.")
		t.Error()
		t.Error("Expected:")
		t.Error(expected)
		t.Error()
		t.Error("Actual:")
		t.Error(actual)
	}
}

func TestNumber(t *testing.T) {
	expected := []string{
		"a", "b",
	}

	AssertTokensEqual(t, expected, Lex(""))
}
