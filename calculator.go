package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	INTEGER  = "integer"
	PLUS     = "plus"
	MINUS    = "minus"
	MULTIPLY = "multiply"
	EOF      = "eof"
)

type Token struct {
	tokenType  string
	tokenValue string
}

func (self *Token) isOperator() bool {
	operators := [4]string{PLUS, MINUS, MULTIPLY}
	for _, operator := range operators {
		if strings.EqualFold(operator, self.tokenType) {
			return true
		}
	}

	return false
}

type Calculator struct {
	currentToken Token
	input        string
	pos          int
}

func (self *Calculator) readInteger() Token {
	str := ""
	for self.pos < len(self.input) && unicode.IsDigit(rune(self.input[self.pos])) {
		str += self.input[self.pos : self.pos+1]
		self.pos++
	}
	return Token{INTEGER, str}
}

func (self *Calculator) readOperator() Token {
	var token Token

	switch self.input[self.pos] {
	case '+':
		token = Token{PLUS, "+"}
	case '-':
		token = Token{MINUS, "-"}
	case '*':
		token = Token{MULTIPLY, "*"}
	default:
		panic("Invalid operator")
	}

	self.pos++
	return token
}

func (self *Calculator) skipSpaces() {
	for unicode.IsSpace(rune(self.input[self.pos])) {
		self.pos++
	}
}

func isOperator(char byte) bool {
	return char == '+' || char == '-' || char == '*'
}

func (self *Calculator) nextToken() Token {
	if self.pos >= len(self.input) {
		self.currentToken = Token{tokenType: EOF}
	} else if unicode.IsDigit(rune(self.input[self.pos])) {
		self.currentToken = self.readInteger()
	} else if isOperator(self.input[self.pos]) {
		self.currentToken = self.readOperator()
	} else if unicode.IsSpace(rune(self.input[self.pos])) {
		self.skipSpaces()
		self.nextToken()
	} else {
		panic("Not recognized input")
	}

	return self.currentToken
}

// Parser

func (self *Calculator) eat(tokenType string) {
	if self.currentToken.tokenType == tokenType {
		self.nextToken()
	} else {
		panic("Token type not matched")
	}
}

func (self *Calculator) intVal() int {
	token := self.currentToken
	self.eat(INTEGER)
	num, err := strconv.Atoi(token.tokenValue)
	if err != nil {
		panic("Not a valid integer")
	}
	return num
}

func binaryCalc(num1 int, num2 int, operator string) int {
	switch operator {
	case PLUS:
		return num1 + num2
	case MINUS:
		return num1 - num2
	case MULTIPLY:
		return num1 * num2
	default:
		panic("Invalid operator")
	}
}

func (self *Calculator) calc(input string) string {
	self.input = input
	self.pos = 0
	self.nextToken()
	result := self.intVal()

	for self.currentToken.tokenType != EOF {
		var operator string

		if self.currentToken.isOperator() {
			operator = self.currentToken.tokenType
		} else {
			panic("Expect operator")
		}

		self.nextToken()
		result = binaryCalc(result, self.intVal(), operator)
	}

	return strconv.Itoa(result)
}

func main() {
	calc := Calculator{}
	fmt.Printf(calc.calc("121  * 301 *3") + "\n")
}
