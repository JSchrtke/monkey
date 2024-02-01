package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	currentChar  byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (self *Lexer) NextToken() token.Token {
	var tok token.Token

	for isWhitespace(self.currentChar) {
		self.readChar()
	}

	switch self.currentChar {
	case '=':
		if self.peek() == '=' {
			self.readChar()
			tok = newToken(token.EQ, "==")
		} else {
			tok = newToken(token.ASSING, string(self.currentChar))
		}
	case '+':
		tok = newToken(token.PLUS, string(self.currentChar))
	case '-':
		tok = newToken(token.MINUS, string(self.currentChar))
	case '/':
		tok = newToken(token.SLASH, string(self.currentChar))
	case '*':
		tok = newToken(token.ASTERISK, string(self.currentChar))
	case '!':
		if self.peek() == '=' {
			self.readChar()
			tok = newToken(token.NEQ, "!=")
		} else {
			tok = newToken(token.BANG, string(self.currentChar))
		}
	case '<':
		tok = newToken(token.LT, string(self.currentChar))
	case '>':
		tok = newToken(token.GT, string(self.currentChar))
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
			return newToken(token.INT, self.readInteger())
		}

		if isLetter(self.currentChar) {
			word := self.readWord()
			if t, ok := token.KeywordFromString(word); ok {
				return t
			}
			return newToken(token.IDENT, word)
		}
	}

	self.readChar()
	return tok
}

func (self *Lexer) peek() byte {
	if self.nextPosition >= len(self.input) {
		return 0
	} else {
		return self.input[self.nextPosition]
	}
}

func (self *Lexer) readChar() {
	if self.nextPosition >= len(self.input) {
		self.currentChar = 0
	} else {
		self.currentChar = self.input[self.nextPosition]
	}
	self.position = self.nextPosition
	self.nextPosition += 1
}

func (self *Lexer) readInteger() string {
	start := self.position
	for isDigit(self.currentChar) {
		self.readChar()
	}
	return self.input[start:self.position]
}

func (self *Lexer) readWord() string {
	start := self.position
	for isLetter(self.currentChar) {
		self.readChar()
	}
	return self.input[start:self.position]
}

func isWhitespace(b byte) bool {
	return b == 9 || b == 32 || b == 10 || b == 13
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func isDigit(b byte) bool {
	return 48 <= b && b <= 57
}

func isLetter(b byte) bool {
	return 65 <= b && b <= 90 || 97 <= b && b <= 122
}
