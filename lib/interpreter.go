package lib

import (
	"strconv"
)

func Interpreter(lexer *lexer) interpreter {
	self := interpreter{lexer: lexer}
	self.currentToken = lexer.nextToken()
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
		panic("Token type not matched")
	}
}

func (self *interpreter) integer() int {
	token := self.currentToken
	self.eat(INTEGER)
	num, err := strconv.Atoi(token.tokenValue)

	if err != nil {
		panic("Not a valid integer")
	}
	return num
}

func (self *interpreter) Factor() int {
	if self.currentToken.tokenType == INTEGER {
		return self.integer()
	} else if self.currentToken.tokenType == LPAREN {
		self.eat(LPAREN)
		result := self.Expr()
		self.eat(RPAREN)
		return result
	} else {
		panic("Invalid token")
	}
}

func (self *interpreter) Term() int {
	result := self.Factor()
	tokenType := self.currentToken.tokenType

	for tokenType == MUL || tokenType == DIV {
		if tokenType == MUL {
			self.eat(MUL)
			result *= self.Factor()
		} else if tokenType == DIV {
			self.eat(DIV)
			result /= self.Factor()
		}

		tokenType = self.currentToken.tokenType
	}

	return result
}

func (self *interpreter) Expr() int {
	result := self.Term()
	tokenType := self.currentToken.tokenType

	for tokenType == PLUS || tokenType == MINUS {
		if tokenType == PLUS {
			self.eat(PLUS)
			result += self.Term()
		} else if tokenType == MINUS {
			self.eat(MINUS)
			result -= self.Term()
		}

		tokenType = self.currentToken.tokenType
	}

	return result
}
