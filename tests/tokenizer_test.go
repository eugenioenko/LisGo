package testing

import (
	"lisgo/pkg/lisgo"
	"testing"
)

func TestTokenizerHandlesNumbers(t *testing.T) {
	tokenizer := lisgo.MakeTokenizer()
	tokenizer.LoadFromString("123 -456 78.90 -0.12")
	tokens := tokenizer.Tokenize()

	expected := []struct {
		Type    lisgo.TokenType
		Literal string
	}{
		{lisgo.TokenTypeInteger, "123"},
		{lisgo.TokenTypeInteger, "-456"},
		{lisgo.TokenTypeFloat, "78.90"},
		{lisgo.TokenTypeFloat, "-0.12"},
	}

	for i, token := range tokens[:len(expected)] {
		if token.Type != expected[i].Type || token.Literal != expected[i].Literal {
			t.Errorf("Expected %v, got %v", expected[i], token)
		}
	}
}

func TestTokenizerHandlesStrings(t *testing.T) {
	tokenizer := lisgo.MakeTokenizer()
	tokenizer.LoadFromString(`"hello" 'world'`)
	tokens := tokenizer.Tokenize()

	expected := []struct {
		Type    lisgo.TokenType
		Literal string
	}{
		{lisgo.TokenTypeString, "hello"},
		{lisgo.TokenTypeString, "world"},
	}

	for i, token := range tokens[:len(expected)] {
		if token.Type != expected[i].Type || token.Literal != expected[i].Literal {
			t.Errorf("Expected %v, got %v", expected[i], token)
		}
	}
}

func TestTokenizerHandlesIdentifiers(t *testing.T) {
	tokenizer := lisgo.MakeTokenizer()
	tokenizer.LoadFromString("variable_name another-variable _underscore")
	tokens := tokenizer.Tokenize()

	expected := []struct {
		Type    lisgo.TokenType
		Literal string
	}{
		{lisgo.TokenTypeIdentifier, "variable_name"},
		{lisgo.TokenTypeIdentifier, "another-variable"},
		{lisgo.TokenTypeIdentifier, "_underscore"},
	}

	for i, token := range tokens[:len(expected)] {
		if token.Type != expected[i].Type || token.Literal != expected[i].Literal {
			t.Errorf("Expected %v, got %v", expected[i], token)
		}
	}
}

func TestTokenizerHandlesComments(t *testing.T) {
	tokenizer := lisgo.MakeTokenizer()
	tokenizer.LoadFromString("123 ; this is a comment\n456")
	tokens := tokenizer.Tokenize()

	expected := []struct {
		Type    lisgo.TokenType
		Literal string
	}{
		{lisgo.TokenTypeInteger, "123"},
		{lisgo.TokenTypeInteger, "456"},
	}

	for i, token := range tokens[:len(expected)] {
		if token.Type != expected[i].Type || token.Literal != expected[i].Literal {
			t.Errorf("Expected %v, got %v", expected[i], token)
		}
	}
}

func TestTokenizerHandlesParentheses(t *testing.T) {
	tokenizer := lisgo.MakeTokenizer()
	tokenizer.LoadFromString("( )")
	tokens := tokenizer.Tokenize()

	expected := []struct {
		Type    lisgo.TokenType
		Literal string
	}{
		{lisgo.TokenTypeLeftParen, "("},
		{lisgo.TokenTypeRightParen, ")"},
	}

	for i, token := range tokens[:len(expected)] {
		if token.Type != expected[i].Type || token.Literal != expected[i].Literal {
			t.Errorf("Expected %v, got %v", expected[i], token)
		}
	}
}

func TestTokenizerHandlesReservedTokens(t *testing.T) {
	tokenizer := lisgo.MakeTokenizer()
	tokenizer.LoadFromString("if while true false null")
	tokens := tokenizer.Tokenize()

	expected := []struct {
		Type    lisgo.TokenType
		Literal string
	}{
		{lisgo.TokenTypeReserved, "if"},
		{lisgo.TokenTypeReserved, "while"},
		{lisgo.TokenTypeTrue, "true"},
		{lisgo.TokenTypeFalse, "false"},
		{lisgo.TokenTypeNull, "null"},
	}

	for i, token := range tokens[:len(expected)] {
		if token.Type != expected[i].Type || token.Literal != expected[i].Literal {
			t.Errorf("Expected %v, got %v", expected[i], token)
		}
	}
}

func TestTokenizerHandlesUnexpectedCharacters(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected tokenizer to panic on unexpected character")
		}
	}()

	tokenizer := lisgo.MakeTokenizer()
	tokenizer.LoadFromString("@")
	tokenizer.Tokenize()
}
