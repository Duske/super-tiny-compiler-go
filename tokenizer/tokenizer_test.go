package tokenizer

import (
	"testing"
	"reflect"
)

func TestTokenizer(t *testing.T) {
	const input = "(add 2 (subtract 4 2))"
	correctTokens := []Token{
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
	builtTokens := Build(input)
	if len(builtTokens) != len(correctTokens) {
		t.Log("The number of tokens is not correct")
		t.Fail()
	}
	if !reflect.DeepEqual(builtTokens, correctTokens) {
		t.Log("Tokens are not equal to target output")
		t.Fail()
	}
}