package main

import (
	"fmt"
	"os"
	"super-tiny-compiler/tokenizer"
	"super-tiny-compiler/parser"
)

func main() {
	fmt.Println("Starting tokenizer")
	if len(os.Args) > 1 {
		tokens := tokenizer.Build(os.Args[1])
		fmt.Println(tokens)
		ast := parser.Parse(tokens)
		fmt.Println(ast)
	} else {
		fmt.Printf("No string provided")
	}
	fmt.Println("End tokenizer")
}