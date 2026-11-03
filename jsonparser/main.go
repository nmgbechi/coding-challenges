package main

import (
	"fmt"

	"github.com/nmgbechi/coding-challenges/jsonparser/lexer"
	"github.com/nmgbechi/coding-challenges/jsonparser/token"
)

func main() {
	input := "{ [ ,                                         ] }                  "

	lex := lexer.New(input)

	curr := lex.NextToken()
	for curr.TokenType != token.EOF {
		fmt.Println(curr)
		curr = lex.NextToken()
	}
}
