## Golang calculator

An implementation of an arithmetic calculator in Golang. This implementation follows the structure in [Let's build a simple interpreter](http://ruslanspivak.com/lsbasi-part6/). Here's the reference grammar:

    expr:   muldiv ((PLUS | MINUS) muldiv)*
    term:   factor ((MUL | DIV) factor)*
    factor: INTEGER | LPAREN expr RPAREN

Running the interactive calculator:

    go build
    ./go-calculator

or just do `go run main.go`.

Some example calculations:

    > 1+2
    3
    > (12 + 1) / (1 + 2)
    4.333333333
    > 3 + 2*4/8
    4
    > (1 - (3 + 2) * 42) / 5
    -41.8
    > 1a + 1
    Invalid token 'a'
    > 1 +
    Unexpected token '{eof }'
    > ((1-3 * 2)
    Token '{eof }' does not match expected type 'right parenthesis'
