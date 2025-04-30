package testing

import (
	"lisgo/pkg/lisgo"
	"testing"
)

func TestParserHandlesEmptyInput(t *testing.T) {
	parser := lisgo.MakeParser()
	tokens := []lisgo.Token{}
	expressions := parser.Parse(tokens)
	if len(expressions) != 0 {
		t.Errorf("Expected no expressions, got %d", len(expressions))
	}
}

func TestParserHandlesUnmatchedParentheses(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected parser to panic on unmatched parentheses")
		}
	}()
	parser := lisgo.MakeParser()
	tokens := []lisgo.Token{
		{Type: lisgo.TokenTypeLeftParen, Literal: "("},
		{Type: lisgo.TokenTypeIdentifier, Literal: "test"},
	}
	parser.Parse(tokens)
}
