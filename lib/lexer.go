package lib

import (
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
	return lexer{input, 0}
}

type lexer struct {
	input string
	pos   int
}

func (self *lexer) intToken() Token {
	str := ""
	for self.pos < len(self.input) && unicode.IsDigit(rune(self.input[self.pos])) {
		str += self.input[self.pos : self.pos+1]
		self.pos++
	}
	return Token{INTEGER, str}
}

func (self *lexer) opToken() Token {
	var token Token

	switch self.input[self.pos] {
	case '+':
		token = Token{PLUS, "+"}
	case '-':
		token = Token{MINUS, "-"}
	case '*':
		token = Token{MUL, "*"}
	case '/':
		token = Token{DIV, "/"}
	default:
		panic("Invalid operator")
	}

	self.pos++
	return token
}

func (self *lexer) parenToken() Token {
	var token Token

	switch self.input[self.pos] {
	case '(':
		token = Token{LPAREN, "("}
	case ')':
		token = Token{RPAREN, ")"}
	}

	self.pos++
	return token
}

func (self *lexer) skipSpaces() {
	for unicode.IsSpace(rune(self.input[self.pos])) {
		self.pos++
	}
}

func (self *lexer) nextToken() Token {
	if self.pos >= len(self.input) {
		return Token{tokenType: EOF}
	} else if unicode.IsDigit(rune(self.input[self.pos])) {
		return self.intToken()
	} else if isOperator(self.input[self.pos]) {
		return self.opToken()
	} else if isParenthesis(self.input[self.pos]) {
		return self.parenToken()
	} else if unicode.IsSpace(rune(self.input[self.pos])) {
		self.skipSpaces()
		return self.nextToken()
	} else {
		panic("Not recognized input")
	}
}
