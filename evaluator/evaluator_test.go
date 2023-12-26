package evaluator

import (
	"testing"

	"github.com/Captainmango/monkey/lexer"
	"github.com/Captainmango/monkey/object"
	"github.com/Captainmango/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct{
		input string
		expected int64
	} {
		{
			"5", 
			5,
		},
		{
			"10", 
			10,
		},
	}

	for _, tc := range tests {
		evaluated := testEval(tc.input)
		testIntegerObj(t, evaluated, tc.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	return Eval(program)
}

func testIntegerObj(t testing.TB, obj object.Object, want int64) bool {
	result, ok := obj.(*object.Integer)

	if !ok {
		t.Errorf("object is not an integer. Got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != want {
		t.Errorf("object has incorrect value. Want %d, got %d", want, result.Value)
		return false
	}

	return true
}