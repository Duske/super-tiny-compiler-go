package traverser

import (
	"github.com/duske/super-tiny-compiler-go/parser"
	"reflect"
)

type Visitor map[string]func(n parser.ASTType, p parser.ASTType)

func traverseNode(nodeType parser.ASTType, parent parser.ASTType, visitor Visitor) {
	switch node := nodeType.(type) {
	case parser.NumberLiteral, parser.StringLiteral:
		visitor[reflect.TypeOf(node).Name()](node, parent)
		break
	case parser.CallExpression:
		visitor[reflect.TypeOf(node).Name()](node, parent)
		traverseArray(node.Params, parent, visitor)
		break
	}
}

func traverseArray(nodes []parser.ASTType, parent parser.ASTType, visitor Visitor) {
	for _, node := range nodes {
		traverseNode(node, parent, visitor)
	}
}

func Traverser(ast parser.AST, visitor Visitor) {
	traverseArray(ast.Body, nil, visitor)
}
