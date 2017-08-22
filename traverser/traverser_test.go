package traverser

import (
	"github.com/duske/super-tiny-compiler-go/parser"
	"testing"
)

func TestTraverser(t *testing.T) {
	enteredNums := 0
	enteredStrings := 0
	enteredCallexp := 0
	v := Visitor{
		"NumberLiteral": func(n parser.ASTType, p parser.ASTType) {
			enteredNums = enteredNums + 1
		},
		"StringLiteral": func(n parser.ASTType, p parser.ASTType) {
			enteredStrings = enteredStrings + 1
		},
		"CallExpression": func(n parser.ASTType, p parser.ASTType) {
			enteredCallexp = enteredCallexp + 1
		},
	}
	correctAst := parser.AST{TypeName: "Program", Body: []parser.ASTType{
		parser.CallExpression{Value: "add", Params: []parser.ASTType{
			parser.NumberLiteral{Value: "2"},
			parser.CallExpression{Value: "subtract", Params: []parser.ASTType{
				parser.NumberLiteral{Value: "4"},
				parser.NumberLiteral{Value: "2"},
			}},
		}},
	}}
	Traverser(correctAst, v)
	if enteredNums != 3 {
		t.Log("NumberLiteral was not visited correctly")
		t.Fail()
	}
	if enteredStrings != 0 {
		t.Log("StringLiteral was not visited correctly")
		t.Fail()
	}
	if enteredCallexp != 2 {
		t.Log("CallExpression was not visited correctly")
		t.Fail()
	}
}
