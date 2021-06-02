package lexer

import (
	"ceqi/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, wanted: %q, got: %q)", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, wanted: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestSkipWhitespace(t *testing.T) {
	l := New(`

	let the_game_begain`)

	l.skipWhitespace()

	got := l.ch

	if got != 'l' {
		t.Fatalf("l.skipWhitespace(), l.ch = %v, expected: l", got)
	}

}
