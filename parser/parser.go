package parser

import (
	"github.com/duske/super-tiny-compiler-go/tokenizer"
)

type ASTType interface {
	Value() string
}

func (node NumberLiteral) Value() string {
	return node.value
}

func (node StringLiteral) Value() string {
	return node.value
}

func (node CallExpression) Value() string {
	return node.value
}

type NumberLiteral struct {
	value string
}

type StringLiteral struct {
	value string
}

type CallExpression struct {
	value  string
	params []ASTType
}

type AST struct {
	typeName string
	body     []ASTType
}

func Walk(tokensP *[]tokenizer.Token, currentP *int) ASTType {
	var token tokenizer.Token
	tokens := *tokensP
	token = tokens[*currentP]

	if token.TypeName == "number" {
		*currentP++

		return NumberLiteral{value: token.Value}
	}

	if token.TypeName == "string" {
		*currentP++

		return StringLiteral{value: token.Value}
	}
	if token.TypeName == "paren" && token.Value == "(" {
		// go to next token
		*currentP++
		token = tokens[*currentP]
		// use token's value for callexpression
		callExpression := CallExpression{value: token.Value}

		// go to next token
		*currentP++
		token = tokens[*currentP]

		for token.TypeName != "paren" || (token.TypeName == "paren" && token.Value != ")") {
			callExpression.params = append(callExpression.params, Walk(&tokens, currentP))
			// continue with new token position from the terminated recursive call of walk()
			token = tokens[*currentP]
		}
		*currentP++
		return callExpression
	}

	return nil

}

func Parse(tokens []tokenizer.Token) AST {
	current := 0
	ast := AST{typeName: "Program"}
	for current < len(tokens)-1 {
		ast.body = append(ast.body, Walk(&tokens, &current))
	}
	return ast
}
