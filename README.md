## Golang calculator

An implementation of an arithmetic calculator in Golang. This implementation follows the structure in [Let's build a simple interpreter](http://ruslanspivak.com/lsbasi-part6/). Here's the reference grammar:

    expr:   muldiv ((PLUS | MINUS) muldiv)*
    term:   factor ((MUL | DIV) factor)*
    factor: INTEGER | LPAREN expr RPAREN

Running the interactive calculator:

    go build
    ./go-calculator

or just do `go run main.go`.
