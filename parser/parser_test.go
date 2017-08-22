package parser

import (
	"github.com/duske/super-tiny-compiler-go/tokenizer"
	"reflect"
	"testing"
)

func TestParserWithEmptyTokens(t *testing.T) {
	input := []tokenizer.Token{}
	correctAST := AST{TypeName: "Program"}
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
	correctAST := AST{TypeName: "Program", Body: []ASTType{
		CallExpression{Value: "add", Params: []ASTType{
			NumberLiteral{"2"},
			CallExpression{Value: "subtract", Params: []ASTType{
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
