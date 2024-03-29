package lexer

import (
	"monkey/token"
)

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

	self.skipWhitespace()

	switch self.currentChar {
	case '=':
		tok = newToken(token.ASSING, string(self.currentChar))
	case '+':
		tok = newToken(token.PLUS, string(self.currentChar))
	case ',':
		tok = newToken(token.COMMA, string(self.currentChar))
	case ';':
		tok = newToken(token.SEMICOLON, string(self.currentChar))
	case '(':
		tok = newToken(token.LPAREN, string(self.currentChar))
	case ')':
		tok = newToken(token.RPAREN, string(self.currentChar))
	case '{':
		tok = newToken(token.LBRACE, string(self.currentChar))
	case '}':
		tok = newToken(token.RBRACE, string(self.currentChar))
	case 0:
		tok = newToken(token.EOF, "")
	default:
		if isDigit(self.currentChar) {
			tok = newToken(token.INT, self.readNumber())
			return tok
		}

		if isLetter(self.currentChar) {
			word := self.readLetters()

			if word == "let" {
				tok = newToken(token.LET, word)
				return tok
			}

			if word == "fn" {
				tok = newToken(token.FUNCTION, word)
				return tok
			}

			tok = newToken(token.IDENT, word)
			return tok
		}
	}

	self.readChar()
	return tok
}

func (self *Lexer) skipWhitespace() {
	for isWhitespace(self.currentChar) {
		self.readChar()
	}
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

// is the vertical tab a thing that needs to be considered?
func isWhitespace(b byte) bool {
	return b == 9 || b == 32 || b == 10 || b == 13
}

func (self *Lexer) readNumber() string {
	digits := ""
	for isDigit(self.currentChar) {
		digits += string(self.currentChar)
		self.readChar()
	}
	return digits
}

func isDigit(b byte) bool {
	return 48 <= b && b <= 57
}

func isLetter(b byte) bool {
	return 65 <= b && b <= 90 || 97 <= b && b <= 122
}

func (self *Lexer) readLetters() string {
	letters := ""
	for isLetter(self.currentChar) {
		letters = letters + string(self.currentChar)
		self.readChar()
	}
	return letters
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
