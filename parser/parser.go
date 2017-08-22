package parser

import (
	"github.com/duske/super-tiny-compiler-go/tokenizer"
)

type ASTType interface {
	NodeValue() string
}

func (node NumberLiteral) NodeValue() string {
	return node.Value
}

func (node StringLiteral) NodeValue() string {
	return node.Value
}

func (node CallExpression) NodeValue() string {
	return node.Value
}

type NumberLiteral struct {
	Value string
}

type StringLiteral struct {
	Value string
}

type CallExpression struct {
	Value  string
	Params []ASTType
}

type AST struct {
	TypeName string
	Body     []ASTType
}

func Walk(tokensP *[]tokenizer.Token, currentP *int) ASTType {
	var token tokenizer.Token
	tokens := *tokensP
	token = tokens[*currentP]

	if token.TypeName == "number" {
		*currentP++

		return NumberLiteral{Value: token.Value}
	}

	if token.TypeName == "string" {
		*currentP++

		return StringLiteral{Value: token.Value}
	}
	if token.TypeName == "paren" && token.Value == "(" {
		// go to next token
		*currentP++
		token = tokens[*currentP]
		// use token's value for callexpression
		callExpression := CallExpression{Value: token.Value}

		// go to next token
		*currentP++
		token = tokens[*currentP]

		for token.TypeName != "paren" || (token.TypeName == "paren" && token.Value != ")") {
			callExpression.Params = append(callExpression.Params, Walk(&tokens, currentP))
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
	ast := AST{TypeName: "Program"}
	for current < len(tokens)-1 {
		ast.Body = append(ast.Body, Walk(&tokens, &current))
	}
	return ast
}
