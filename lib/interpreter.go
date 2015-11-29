package lib

import (
	"errors"
	"fmt"
	"strconv"
)

func Interpreter(lexer *lexer) interpreter {
	self := interpreter{lexer: lexer}
	return self
}

type interpreter struct {
	currentToken Token
	lexer        *lexer
}

func (self *interpreter) eat(tokenType string) {
	if self.currentToken.tokenType == tokenType {
		self.currentToken = self.lexer.nextToken()
	} else {
		panic(fmt.Sprintf("Token '%v' does not match expected type '%s'", self.currentToken, tokenType))
	}
}

func (self *interpreter) integer() int {
	token := self.currentToken
	self.eat(INTEGER)
	num, err := strconv.Atoi(token.tokenValue)

	if err != nil {
		panic(fmt.Sprintf("'%s' is not a valid integer", token.tokenValue))
	}
	return num
}

func (self *interpreter) factor() int {
	if self.currentToken.tokenType == INTEGER {
		return self.integer()
	} else if self.currentToken.tokenType == LPAREN {
		self.eat(LPAREN)
		result := self.expr()
		self.eat(RPAREN)
		return result
	} else {
		panic(fmt.Sprintf("Unexpected token '%v'", self.currentToken))
	}
}

func (self *interpreter) term() int {
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

func (self *interpreter) expr() int {
	result := self.term()
	tokenType := self.currentToken.tokenType

	for tokenType == PLUS || tokenType == MINUS {
		if tokenType == PLUS {
			self.eat(PLUS)
			result += self.term()
		} else if tokenType == MINUS {
			self.eat(MINUS)
			result -= self.term()
		}

		tokenType = self.currentToken.tokenType
	}

	return result
}

func (self *interpreter) Result() (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			result = 0
			err = errors.New(e.(string))
		}
	}()

	self.currentToken = self.lexer.nextToken()
	result = self.expr()
	err = nil
	return result, err
}
