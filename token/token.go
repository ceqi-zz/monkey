package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

//TokenType constants
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers
	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
