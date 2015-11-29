package lib

import (
	"fmt"
	"unicode"
)

const (
	INTEGER = "integer"
	PLUS    = "plus"
	MINUS   = "minus"
	MUL     = "mul"
	DIV     = "div"
	EOF     = "eof"
	LPAREN  = "left parenthesis"
	RPAREN  = "right parenthesis"
)

func isOperator(char byte) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func isParenthesis(char byte) bool {
	return char == '(' || char == ')'
}

type Token struct {
	tokenType  string
	tokenValue string
}

func Lexer(input string) lexer {
	return lexer{input, 0, input[0]}
}

type lexer struct {
	input       string
	pos         int
	currentChar byte
}

func (self *lexer) advance() {
	self.pos++
	if self.pos < len(self.input) {
		self.currentChar = self.input[self.pos]
	} else {
		self.currentChar = '\x00'
	}
}

func (self *lexer) intToken() Token {
	str := ""
	for self.pos < len(self.input) && unicode.IsDigit(rune(self.currentChar)) {
		str += self.input[self.pos : self.pos+1]
		self.advance()
	}
	return Token{INTEGER, str}
}

func (self *lexer) opToken() Token {
	var token Token

	switch self.currentChar {
	case '+':
		token = Token{PLUS, "+"}
	case '-':
		token = Token{MINUS, "-"}
	case '*':
		token = Token{MUL, "*"}
	case '/':
		token = Token{DIV, "/"}
	default:
		panic(fmt.Sprintf("Invalid operator '%c'", self.currentChar))
	}

	self.advance()
	return token
}

func (self *lexer) parenToken() Token {
	var token Token

	switch self.currentChar {
	case '(':
		token = Token{LPAREN, "("}
	case ')':
		token = Token{RPAREN, ")"}
	}

	self.advance()
	return token
}

func (self *lexer) skipSpaces() {
	for self.pos < len(self.input) && unicode.IsSpace(rune(self.currentChar)) {
		self.advance()
	}
}

func (self *lexer) nextToken() Token {
	if self.pos >= len(self.input) {
		return Token{tokenType: EOF}
	} else if unicode.IsDigit(rune(self.currentChar)) {
		return self.intToken()
	} else if isOperator(self.currentChar) {
		return self.opToken()
	} else if isParenthesis(self.currentChar) {
		return self.parenToken()
	} else if unicode.IsSpace(rune(self.currentChar)) {
		self.skipSpaces()
		return self.nextToken()
	} else {
		panic(fmt.Sprintf("Invalid token '%c'", self.currentChar))
	}
}
