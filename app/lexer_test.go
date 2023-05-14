package app

import (
	"testing"
)

func areTokensEqual(expected []string, actual []string) bool {
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

func assertTokensEqual(t *testing.T, json string, expected []string) {
	actual := Lex(json)

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
	assertTokensEqual(t, "1", []string{"1"})
	assertTokensEqual(t, "12", []string{"12"})
	assertTokensEqual(t, "123", []string{"123"})

	assertTokensEqual(t, "1x3", []string{})
}

func TestBoolean(t *testing.T) {
	assertTokensEqual(t, "false", []string{"false"})
	assertTokensEqual(t, "true", []string{"true"})
}

func TestString(t *testing.T) {
	assertTokensEqual(t, `"abc"`, []string{"abc"})
}
