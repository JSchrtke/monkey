package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	5 != 10;
	`

	expectedTokens := []token.Token{
		// let five = 5;
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.ASSING, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},

		// let ten = 10;
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.ASSING, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},

		// let add = fn(x, y) {
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.ASSING, Literal: "="},
		{Type: token.FUNCTION, Literal: "fn"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},

		// 	x + y;
		{Type: token.IDENT, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},

		// };
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.SEMICOLON, Literal: ";"},

		// let result = add(five, ten);
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "result"},
		{Type: token.ASSING, Literal: "="},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},

		// !-/*5;
		{Type: token.BANG, Literal: "!"},
		{Type: token.MINUS, Literal: "-"},
		{Type: token.SLASH, Literal: "/"},
		{Type: token.ASTERISK, Literal: "*"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},

		// 5 < 10 > 5;
		{Type: token.INT, Literal: "5"},
		{Type: token.LT, Literal: "<"},
		{Type: token.INT, Literal: "10"},
		{Type: token.GT, Literal: ">"},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},

		// if (5 < 10) {
		{Type: token.IF, Literal: "if"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.INT, Literal: "5"},
		{Type: token.LT, Literal: "<"},
		{Type: token.INT, Literal: "10"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},

		// 	return true;
		{Type: token.RETURN, Literal: "return"},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},

		// } else {
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.ELSE, Literal: "else"},
		{Type: token.LBRACE, Literal: "{"},

		// 	return false;
		{Type: token.RETURN, Literal: "return"},
		{Type: token.FALSE, Literal: "false"},
		{Type: token.SEMICOLON, Literal: ";"},

		// }
		{Type: token.RBRACE, Literal: "}"},

		// 10 == 10;
		{Type: token.INT, Literal: "10"},
		{Type: token.EQ, Literal: "=="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},

		// 5 != 10;
		{Type: token.INT, Literal: "5"},
		{Type: token.NEQ, Literal: "!="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.EOF, Literal: ""},
	}

	lexer := New(input)

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
