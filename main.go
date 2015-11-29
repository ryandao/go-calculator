package main

import (
	"fmt"
	"github.com/ryandao/calculator/lib"
)

func main() {
	lexer := lib.Lexer("2 * (2  + 10) *3/3 + 1")
	interpreter := lib.Interpreter(&lexer)
	fmt.Printf("%d\n", interpreter.Expr())
}
