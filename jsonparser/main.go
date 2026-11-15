package main

import (
	"fmt"

	"github.com/nmgbechi/coding-challenges/jsonparser/lexer"
	"github.com/nmgbechi/coding-challenges/jsonparser/token"
)

func main() {
	input := "{ [ ,                                         ] }                  "
	input = `{"name": "Bob", "age": 42, "ok": true}`
	input = `"hi" "there" ""`
	lex := lexer.New(input)

	curr := lex.NextToken()
	for curr.TokenType != token.EOF {
		fmt.Println(curr)
		curr = lex.NextToken()
	}

}
