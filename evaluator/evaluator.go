package evaluator

import (
	"github.com/Captainmango/monkey/ast"
	"github.com/Captainmango/monkey/object"
)

var (
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBoolObj(node.Value)
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, st := range stmts {
		result = Eval(st)
	}

	return result
}

func nativeBoolToBoolObj(input bool) *object.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}