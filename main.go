package main

import (
	"fmt"
	"github.com/duske/super-tiny-compiler-go/parser"
	"github.com/duske/super-tiny-compiler-go/tokenizer"
	"os"
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
