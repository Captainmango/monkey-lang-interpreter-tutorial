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

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}

}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)

	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
		result.Value, expected)

		return false
	}
	
	return true
}