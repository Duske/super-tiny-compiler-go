package parser

import (
	"reflect"
	"super-tiny-compiler-go/tokenizer"
	"testing"
)

func TestParserWithEmptyTokens(t *testing.T) {
	input := []tokenizer.Token{}
	correctAST := AST{typeName: "Program"}
	output := Parse(input)
	if !reflect.DeepEqual(correctAST, output) {
		t.Log("Empty tokens does not create an correct AST")
		t.Fail()
	}
}

func TestParserWithTokens(t *testing.T) {
	input := []tokenizer.Token{
		{TypeName: "paren", Value: "("},
		{TypeName: "name", Value: "add"},
		{TypeName: "number", Value: "2"},
		{TypeName: "paren", Value: "("},
		{TypeName: "name", Value: "subtract"},
		{TypeName: "number", Value: "4"},
		{TypeName: "number", Value: "2"},
		{TypeName: "paren", Value: ")"},
		{TypeName: "paren", Value: ")"},
	}
	correctAST := AST{typeName: "Program", body: []ASTType{
		CallExpression{value: "add", params: []ASTType{
			NumberLiteral{"2"},
			CallExpression{value: "subtract", params: []ASTType{
				NumberLiteral{"4"},
				NumberLiteral{"2"},
			}},
		}},
	}}
	output := Parse(input)
	if !reflect.DeepEqual(correctAST, output) {
		t.Log("AST is not correct")
		t.Fail()
	}
}
