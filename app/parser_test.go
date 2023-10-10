package app

import (
	"testing"
)

func assertParsingFails(t *testing.T, json string) {
	_, ok := Lex(json)

	if ok {
		t.Fatal("Expected lexing to fail, but it succeeded.")
	}
}

func areAstsEqual(expected *Node, actual *Node) bool {
	// TODO
	return true
}

func assertAstsEqual(t *testing.T, json string, expected *Node) {
	tokens, lexOk := Lex(json)

	if !lexOk {
		t.Fatal("Lexing failed.")
	}

	actual, parseOk := Parse(tokens)

	if !parseOk {
		t.Fatal("Parsing failed.")
	}

	if !areAstsEqual(expected, actual) {
		// If this happens, use the debugger to see how `expected` and `actual`
		// differ
		t.Fatal("ASTs are not equal.")
	}
}

func TestNumber(t *testing.T) {
	var expected Node = NumberNode{12}

	assertAstsEqual(t, "12", &expected)
}
