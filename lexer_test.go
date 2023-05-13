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

func AssertTokensEqual(t *testing.T, json string, expected []string) {
	actual := Lex(json)

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

func TestInteger(t *testing.T) {
	AssertTokensEqual(t, "1", []string{"1"})
}

// func TestBoolean(t *testing.T) {
// 	AssertTokensEqual(t, "false", []string{"false"})
// 	AssertTokensEqual(t, "true", []string{"true"})
// }
