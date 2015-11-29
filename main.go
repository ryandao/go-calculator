package main

import (
	"bufio"
	"fmt"
	"github.com/ryandao/go-calculator/lib"
	"os"
)

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		lexer := lib.Lexer(input)
		interpreter := lib.Interpreter(&lexer)
		result, err := interpreter.Result()

		if err == nil {
			fmt.Printf("%d\n", result)
		} else {
			fmt.Printf("%s\n", err)
		}
	}
}
