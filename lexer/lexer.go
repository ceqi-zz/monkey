package lexer

import (
	"ceqi/monkey/token"
	"regexp"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, string(l.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	// case '(':
	// 	tok = newToken(token.LPAREN, l.ch)
	// case ')':
	// 	tok = newToken(token.RPAREN, l.ch)
	// case ',':
	// 	tok = newToken(token.COMMA, l.ch)
	// case '+':
	// 	tok = newToken(token.PLUS, l.ch)
	// case '{':
	// 	tok = newToken(token.LBRACE, l.ch)
	// case '}':
	// 	tok = newToken(token.RBRACE, l.ch)
	// case 0:
	// 	tok.Literal = ""
	// 	tok.Type = token.EOF
	default:
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	var identifier []byte
	for isLetter(l.ch) {
		identifier = append(identifier, l.ch) // can use lexer struct instead
		l.readChar()
	}
	return string(identifier)
}

func isLetter(ch byte) bool {
	r := regexp.MustCompile(`[a-zA-Z_]{1,1}`) // can use byte comparison
	return r.Match([]byte{ch})
}

func newToken(tt token.TokenType, tok string) token.Token {
	return token.Token{Type: tt, Literal: tok}
}

func (l *Lexer) skipWhitespace() {
	r := regexp.MustCompile(`\s`)
	for r.Match([]byte{l.ch}) {
		l.readChar()
	}
}
