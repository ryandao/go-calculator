package main

import (
	"bufio"
	"fmt"
	"github.com/ryandao/go-calculator/lib"
	"os"
)

// Strip trailing zeros from a float
func formatFloat(num float64) string {
	str := fmt.Sprintf("%.9f", num)
	truncate := len(str)

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '0' {
			truncate = i
		}
	}

	if truncate > 0 && str[truncate-1] == '.' {
		truncate--
	}

	return str[0:truncate]
}

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		lexer := lib.Lexer(input)
		interpreter := lib.Interpreter(&lexer)
		result, err := interpreter.Result()

		if err == nil {
			fmt.Printf("%s\n", formatFloat(result))
		} else {
			fmt.Printf("%s\n", err)
		}
	}
}
