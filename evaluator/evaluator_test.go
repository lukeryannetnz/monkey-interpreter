package evaluator

import (
	"monkey-interpreter/lexer"
	"monkey-interpreter/object"
	"monkey-interpreter/parser"
	"testing"
)

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
		if(10 > 1) {
			if (10 > 1) {
				return 10;
			}

			return 1;
		}
		`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 - 50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"50 / 2 * 2 + 10", 60},
		{"(2 + 3) * 7", 35},
		{"(100 * 2) / 20", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"true == true", true},
		{"true != true", false},
		{"false == false", true},
		{"false != false", false},
		{"(1 < 2) == true", true},
		{"(1 < 2) != true", false},
		{"(1 < 2) == false", false},
		{"(1 < 2) != false", true},
		{"3 > 9", false},
		{"1 > 2", false},
		{"3 < 9", true},
		{"1 < 2", true},
		{"13 == 13", true},
		{"13 != 13", false},
		{"1 == 2", false},
		{"1 != 2", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestNullExpression(t *testing.T) {
	obj := testEval("")

	result, ok := obj.(*object.Null)

	if !ok {
		t.Errorf("wrong type, expected object.Null. got %T", result.Type())
	}
}

func TestBangExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestAdditionExpression(t *testing.T) {
	obj := testEval("1 + 2")
	result, ok := obj.(*object.Integer)

	if !ok {
		t.Errorf("wrong type, expected object.Integer. got %T", result.Type())
		return
	}

	if result.Value != 3 {
		t.Error("wrong value. expected 3.")
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got = %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("wrong type, expected object.Boolean. got:%T", obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. expected %t, got %t", expected, result.Value)
		return false
	}

	return true
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
				if(10 > 1) {
					if(10 > 1) {
						return true + false;
					}
				}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)

		if !ok {
			t.Errorf("no error object returned. got =%T(%+v) expected error message=%q", evaluated, evaluated, tt.expectedMessage)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}
