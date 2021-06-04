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

	l.eatWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			tok.Literal = l.readDigit()
			tok.Type = token.INT
			return tok
		} else if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) readDigit() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
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

func newToken(tt token.TokenType, tok byte) token.Token {
	return token.Token{Type: tt, Literal: string(tok)}
}

func (l *Lexer) eatWhitespace() {
	r := regexp.MustCompile(`\s`)
	for r.Match([]byte{l.ch}) {
		l.readChar()
	}
}
