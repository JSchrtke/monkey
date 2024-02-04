package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func MakeLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (self *Lexer) NextToken() token.Token {
	var tok token.Token

	switch self.currentChar {
	case '=':
		tok = newToken(token.ASSING, self.currentChar)
	case '+':
		tok = newToken(token.PLUS, self.currentChar)
	case ',':
		tok = newToken(token.COMMA, self.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, self.currentChar)
	case '(':
		tok = newToken(token.LPAREN, self.currentChar)
	case ')':
		tok = newToken(token.RPAREN, self.currentChar)
	case '{':
		tok = newToken(token.LBRACE, self.currentChar)
	case '}':
		tok = newToken(token.RBRACE, self.currentChar)
	case 0:
		tok = newToken(token.EOF, 0)
	default:
		if isLetter(self.currentChar) {
			// todo
		} else {
			tok = newToken(token.ILLEGAL, self.currentChar)
		}
	}

	self.readChar()
	return tok
}

func isLetter(b byte) bool {
	return 65 <= b && b <= 90 || 97 <= b && b <= 122
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (self *Lexer) readChar() {
	if self.readPosition >= len(self.input) {
		self.currentChar = 0
	} else {
		self.currentChar = self.input[self.readPosition]
	}
	self.position = self.readPosition
	self.readPosition += 1
}
