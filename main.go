package main

import (
	"fmt"
	"golang/lexer"
)

func main() {
	code := "int x = 5 + 3 - 2 * (1 + 4/3)"

	tokens := lexer.Lex(code)
	lexer.Filter(tokens)

	fmt.Println(lexer.Stringify(tokens))
}
