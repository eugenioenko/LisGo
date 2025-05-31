package testing

import (
	"testing"

	"github.com/eugenioenko/lisgo/pkg/lisgo"
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

func TestParserHandlesInvalidTokens(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected parser to panic on invalid tokens")
		}
	}()
	parser := lisgo.MakeParser()
	tokens := []lisgo.Token{
		{Type: lisgo.TokenTypeIdentifier, Literal: "invalid_token"},
		{Type: lisgo.TokenTypeEof, Literal: ""},
	}
	parser.Parse(tokens)
}

func TestParserHandlesDeeplyNestedParentheses(t *testing.T) {
	parser := lisgo.MakeParser()
	tokens := []lisgo.Token{
		{Type: lisgo.TokenTypeLeftParen, Literal: "("},
		{Type: lisgo.TokenTypeLeftParen, Literal: "("},
		{Type: lisgo.TokenTypeRightParen, Literal: ")"},
		{Type: lisgo.TokenTypeRightParen, Literal: ")"},
	}
	expressions := parser.Parse(tokens)
	if len(expressions) != 1 {
		t.Errorf("Expected 1 expression, got %d", len(expressions))
	}
}
