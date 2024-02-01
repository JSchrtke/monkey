package lexer

import "monkey/token"

type Lexer struct {}

func MakeLexer(input string) *Lexer {
	return nil
}

func (self *Lexer) NextToken() token.Token {
	return token.Token {
		Type: token.ILLEGAL,
		Literal: "",
	}
}
