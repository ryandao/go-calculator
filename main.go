package main

import (
	"bufio"
	"fmt"
	"github.com/ryandao/calculator/lib"
	"os"
)

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		lexer := lib.Lexer(input)
		interpreter := lib.Interpreter(&lexer)
		fmt.Printf("%d\n", interpreter.Expr())
	}
}
