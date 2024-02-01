package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSING   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	BANG     = "!"

	EQ  = "=="
	NEQ = "!="
	LT  = "<"
	GT  = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"else":   ELSE,
}

func KeywordFromString(s string) (Token, bool) {
	t, ok := keywords[s]
	return Token{Type: t, Literal: s}, ok
}
