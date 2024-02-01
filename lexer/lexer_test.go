package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	expectedTokens := []token.Token {
		{Type: token.ASSING, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
	}

	lexer := MakeLexer(input)

	for _, expected := range expectedTokens {
		actual := lexer.NextToken()

		if actual.Type != expected.Type {
			t.Fatalf(
				"wrong token type: expected '%v', but got '%v'",
				expected.Type,
				actual.Type,
			)
		}

		if actual.Literal != expected.Literal {
			t.Fatalf(
				"wrong token literal: expected '%v', but got '%v'",
				expected.Literal,
				actual.Literal,
			)
		}
	}
}
