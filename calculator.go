package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	INTEGER = "integer"
	PLUS    = "plus"
	MINUS   = "minus"
	MUL     = "mul"
	DIV     = "div"
	EOF     = "eof"
	LPAREN  = "("
	RPAREN  = ")"
)

type Token struct {
	tokenType  string
	tokenValue string
}

func (self *Token) isOperator() bool {
	operators := [4]string{PLUS, MINUS, MUL, DIV}
	for _, operator := range operators {
		if strings.EqualFold(operator, self.tokenType) {
			return true
		}
	}

	return false
}

// Lexer

type Lexer struct {
	input string
	pos   int
}

func (self *Lexer) intToken() Token {
	str := ""
	for self.pos < len(self.input) && unicode.IsDigit(rune(self.input[self.pos])) {
		str += self.input[self.pos : self.pos+1]
		self.pos++
	}
	return Token{INTEGER, str}
}

func (self *Lexer) opToken() Token {
	var token Token

	switch self.input[self.pos] {
	case '+':
		token = Token{PLUS, "+"}
	case '-':
		token = Token{MINUS, "-"}
	case '*':
		token = Token{MUL, "*"}
	case '/':
		token = Token{DIV, "*"}
	default:
		panic("Invalid operator")
	}

	self.pos++
	return token
}

func (self *Lexer) skipSpaces() {
	for unicode.IsSpace(rune(self.input[self.pos])) {
		self.pos++
	}
}

func isOperator(char byte) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func (self *Lexer) nextToken() Token {
	if self.pos >= len(self.input) {
		return Token{tokenType: EOF}
	} else if unicode.IsDigit(rune(self.input[self.pos])) {
		return self.intToken()
	} else if isOperator(self.input[self.pos]) {
		return self.opToken()
	} else if self.input[self.pos] == '(' {
		self.pos++
		return Token{LPAREN, "("}
	} else if self.input[self.pos] == ')' {
		self.pos++
		return Token{RPAREN, ")"}
	} else if unicode.IsSpace(rune(self.input[self.pos])) {
		self.skipSpaces()
		return self.nextToken()
	} else {
		panic("Not recognized input")
	}
}

// Parser/interpreter

type Interpreter struct {
	currentToken Token
	lexer        *Lexer
}

func (self *Interpreter) eat(tokenType string) {
	if self.currentToken.tokenType == tokenType {
		self.currentToken = self.lexer.nextToken()
	} else {
		panic("Token type not matched")
	}
}

func (self *Interpreter) integer() int {
	token := self.currentToken
	self.eat(INTEGER)
	num, err := strconv.Atoi(token.tokenValue)

	if err != nil {
		panic("Not a valid integer")
	}
	return num
}

func (self *Interpreter) operator() string {
	if self.currentToken.isOperator() {
		operator := self.currentToken.tokenType
		self.eat(operator)
		return operator
	}

	panic("Expect operator")
}

func (self *Interpreter) factor() int {
	if self.currentToken.tokenType == INTEGER {
		return self.integer()
	} else if self.currentToken.tokenType == LPAREN {
		self.eat(LPAREN)
		result := self.expr()
		self.eat(RPAREN)
		return result
	} else {
		panic("Invalid token")
	}
}

func (self *Interpreter) muldiv() int {
	result := self.factor()
	tokenType := self.currentToken.tokenType

	for tokenType == MUL || tokenType == DIV {
		if tokenType == MUL {
			self.eat(MUL)
			result *= self.factor()
		} else if tokenType == DIV {
			self.eat(DIV)
			result /= self.factor()
		}

		tokenType = self.currentToken.tokenType
	}

	return result
}

func (self *Interpreter) expr() int {
	result := self.muldiv()

	for self.currentToken.tokenType != EOF && self.currentToken.tokenType != RPAREN {
		operator := self.operator()
		result = binaryCalc(result, self.muldiv(), operator)
	}

	return result
}

func binaryCalc(num1 int, num2 int, operator string) int {
	switch operator {
	case PLUS:
		return num1 + num2
	case MINUS:
		return num1 - num2
	case MUL:
		return num1 * num2
	case DIV:
		return num1 / num2
	default:
		panic("Invalid operator")
	}
}

func main() {
	lexer := Lexer{"2 * (2  + 10) *3/3 + 1", 0}
	interpreter := Interpreter{lexer: &lexer}
	interpreter.currentToken = lexer.nextToken()
	fmt.Printf("%d\n", interpreter.expr())
}
