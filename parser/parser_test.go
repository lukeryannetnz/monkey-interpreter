package parser

import (
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	testNoErrors(t, p)

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedValue      int
	}{
		{"x", 5},
		{"y", 10},
		{"foobar", 838383},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmnt, ok := stmt.(*ast.LetStatement)

	if !ok {
		t.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmnt.Name.Value != name {
		t.Errorf("letStmnt.Name.Value not '%s'. got='%s", name, letStmnt.Name.Value)
		return false
	}

	if letStmnt.Name.TokenLiteral() != name {
		t.Errorf("letStmnt.Name.TokenLiteral() not '%s' got '%s'", name, letStmnt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
return 5;
return 10;
return foobar;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	testNoErrors(t, p)

	tests := []struct {
		expectedToken string
		expectedValue interface{}
	}{
		{"return", 5},
		{"return", 10},
		{"return", "foobar"},
	}

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for i, tst := range tests {
		returnStmt, ok := program.Statements[i].(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", returnStmt.Token)
			continue
		}

		if returnStmt.TokenLiteral() != tst.expectedToken {
			t.Errorf("returnStmt.TokenLiteral not 'return'. got=%q", returnStmt.TokenLiteral())
		}

		testLiteralExpression(t, returnStmt.Value, tst.expectedValue)
	}
}

func testNoErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) != 0 {
		t.Errorf("parser has %d errors", len(errors))
		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not an expression statement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Value.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not *ast.Identifier. got=%T", stmt.Value)
	}

	if ident.Value != "foobar" {
		t.Errorf("value not %s, got=%s", input, ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("tokenLiteral not %s, got=%s", input, ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "13;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not an expression statement. got=%T", program.Statements[0])
	}

	if !testLiteralExpression(t, stmt.Value, 13) {
		return
	}
}

func TestStringLiteralExpression(t *testing.T) {
	input := `"this is a long string"`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not an expression statement. got=%T", program.Statements[0])
	}

	if !testString(t, stmt.Value, "this is a long string") {
		return
	}
}

func TestPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
		{"!false", "!", false},
		{"!true", "!", true},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		testNoErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement is not an expression statement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Value.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expression not *ast.PrefixExpression. got=%T", stmt.Value)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}
	}
}

func TestInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true;", true, "==", true},
		{"true != false;", true, "!=", false},
		{"false == false;", false, "==", false},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		testNoErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has incorrect number of statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement is not an expression statement. got=%T", program.Statements[0])
		}

		if !testInfixExpression(t, stmt.Value, tt.leftValue, tt.operator, tt.rightValue) {
			return
		}
	}
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
	operator string, right interface{}) bool {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expression not *ast.InfixExpression. got=%T", exp)
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Fatalf("opExp.Operator is not '%s'. got=%s", operator, opExp.Operator)
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func TestOperatorPrecendence(t *testing.T) {
	tests :=
		[]struct {
			input    string
			expected string
		}{
			{"-a * b", "((-a) * b)"},
			{"!-a", "(!(-a))"},
			{"a + b + c", "((a + b) + c)"},
			{"a + b - c", "((a + b) - c)"},
			{"a * b * c", "((a * b) * c)"},
			{"a * b / c", "((a * b) / c)"},
			{"a + b / c", "(a + (b / c))"},
			{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
			{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
			{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
			{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
			{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
			{"true", "true"},
			{"false", "false"},
			{"3 > 5 == false", "((3 > 5) == false)"},
			{"3 < 5 == true", "((3 < 5) == true)"},
			{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
			{"(5 + 5) * 2", "((5 + 5) * 2)"},
			{"2 / (5 + 5)", "(2 / (5 + 5))"},
			{"-(5 + 5)", "(-(5 + 5))"},
			{"!(true == true)", "(!(true == true))"},
			{"a * [1, 2, 3, 4][b * c] * d", "((a * ([1, 2, 3, 4][(b * c)])) * d)"},
			{"add(a * b[2], b[1], 2 * [1, 2][1])", "add((a * (b[2])),(b[1]),(2 * ([1, 2][1])))"},
		}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		testNoErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected '%s', got '%s'", tt.expected, actual)
		}
	}

}

func TestBooleanExpression(t *testing.T) {
	input := "true;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not an expression statement. got=%T", program.Statements[0])
	}

	if !testLiteralExpression(t, stmt.Value, true) {
		return
	}
}

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", len(program.Statements))
	}

	exp, ok := stmt.Value.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression, got =%T", program.Statements[0])
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statement. got=%T", len(exp.Consequence.Statements))
	}
}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", len(program.Statements))
	}

	exp, ok := stmt.Value.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression, got =%T", program.Statements[0])
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statement. got=%T", len(exp.Consequence.Statements))
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y;}`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statments does not contain 1 statement. Got: %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ExpressionStatement. got=%T", program.Statements[0])
	}

	function, ok := stmt.Value.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Value is not ast.FunctionLiteral. Got=%T", stmt.Value)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function.Parameters is not 2. Got=%d", len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements does not contain 1 statement. Got=%d", len(function.Body.Statements))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function.Body.Statements[0] is not ExpressionStatement. Got=%T", function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.Value, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x){};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z){};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()
		testNoErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		if stmt == nil || stmt.Value == nil {
			t.Error("no statement with value returned")
		}

		fl := stmt.Value.(*ast.FunctionLiteral)

		if len(fl.Parameters) != len(test.expectedParams) {
			t.Errorf("incorrect number of parameters returned. got:%d expected:%d", len(fl.Parameters), len(test.expectedParams))
		}

		for i, param := range fl.Parameters {
			if param.Value != test.expectedParams[i] {
				t.Errorf("param value incorrect. got:%s expected:%s", param.Value, test.expectedParams[i])
			}
		}
	}
}

func TestCallExpression(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5)"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Errorf("expected 1 statement, got:%d", len(program.Statements))
	}

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	if stmt == nil || stmt.Value == nil {
		t.Error("no statement with value returned")
	}

	ce, ok := stmt.Value.(*ast.CallExpression)
	if !ok || ce == nil {
		t.Error("no call expression returned")
	}

	if !testIdentifier(t, ce.Function, "add") {
		return
	}

	if len(ce.Arguments) != 3 {
		t.Fatalf("wrong length of arguments. got=%d", len(ce.Arguments))
	}

	testLiteralExpression(t, ce.Arguments[0], 1)
	testInfixExpression(t, ce.Arguments[1], 2, "*", 3)
	testInfixExpression(t, ce.Arguments[2], 4, "+", 5)
}

func TestCallExpressionStringParam(t *testing.T) {
	input := `add("four")`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Errorf("expected 1 statement, got:%d", len(program.Statements))
	}

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	if stmt == nil || stmt.Value == nil {
		t.Error("no statement with value returned")
	}

	ce, ok := stmt.Value.(*ast.CallExpression)
	if !ok || ce == nil {
		t.Error("no call expression returned")
	}

	if ce.Function.String() != "add" {
		return
	}

	if len(ce.Arguments) != 1 {
		t.Fatalf("wrong length of arguments. got=%d", len(ce.Arguments))
	}

	if ce.Arguments[0].String() != "four" {
		t.Fatalf("argument wrong, expected \"four\" got:%s", ce.Arguments[0])
	}
}

func TestFunctionLiteralCallExpression(t *testing.T) {
	input := "fn(x) { x; }(5)"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Errorf("expected 1 statement, got:%d", len(program.Statements))
	}

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	if stmt == nil || stmt.Value == nil {
		t.Error("no statement with value returned")
	}

	ce, ok := stmt.Value.(*ast.CallExpression)
	if !ok || ce == nil {
		t.Error("no call expression returned")
	}
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testIntegerLiteral(t *testing.T, stmt ast.Expression, input interface{}) bool {
	ident, ok := stmt.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expression not *ast.IntegerLiteral. got=%T", stmt)
		return false
	}

	if ident.Value != input {
		t.Errorf("value not %d, got=%d", input, ident.Value)
		return false
	}

	if ident.TokenLiteral() != fmt.Sprintf("%d", input) {
		t.Errorf("tokenLiteral not %d, got=%s", input, ident.TokenLiteral())
		return false
	}

	return true
}

func testString(t *testing.T, exp ast.Expression, value string) bool {
	str, ok := exp.(*ast.StringLiteral)
	if !ok {
		t.Errorf("exp not *ast.StringLiteral. got=%T", exp)
		return false
	}

	if str.Value != value {
		t.Errorf("str.Value not %s. got=%s", value, str.Value)
		return false
	}

	if str.TokenLiteral() != value {
		t.Errorf("str.TokenLiteral not %s. got=%s", value,
			str.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, stmt ast.Expression, input interface{}) bool {
	ident, ok := stmt.(*ast.Boolean)
	if !ok {
		t.Fatalf("expression not *ast.Boolean. got=%T", stmt)
	}

	if ident.Value != input {
		t.Errorf("value not %t, got=%t", input, ident.Value)
	}

	if ident.TokenLiteral() != fmt.Sprintf("%t", input) {
		t.Errorf("tokenLiteral not %s, got=%s", input, ident.TokenLiteral())
	}

	return true
}

func TestParsingArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Value.(*ast.ArrayLiteral)

	if !ok {
		t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Value)
	}

	testIntegerLiteral(t, array.Elements[0], int64(1))
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)

}

func TestParsingIndexExpressions(t *testing.T) {
	input := "myArray[1 + 1]"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := stmt.Value.(*ast.IndexExpression)

	if !ok {
		t.Fatalf("exp not ast.IndexExpression. got=%T", stmt.Value)
	}

	if !testIdentifier(t, indexExp.Left, "myArray") {
		return
	}

	if !testInfixExpression(t, indexExp.Index, 1, "+", 1) {
		return
	}
}

func TestParsingHashLiteralsStringKeys(t *testing.T) {
	input := `{"one": 1, "two": 2, "three": 3}`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Value.(*ast.HashLiteral)

	if !ok {
		t.Fatalf("exp not ast.HashLiteral. got=%T", stmt.Value)
	}

	expected := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	if len(hash.Pairs) != len(expected) {
		t.Fatalf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
	}

	for key, value := range hash.Pairs {
		literal, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key not ast.StringLiteral. got=%T", key)
		}

		expectedValue := expected[literal.Value]
		testIntegerLiteral(t, value, expectedValue)
	}
}

func TestParsingEmptyHashLiteral(t *testing.T) {
	input := "{}"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Value.(*ast.HashLiteral)

	if !ok {
		t.Fatalf("exp not ast.HashLiteral. got=%T", stmt.Value)
	}

	if len(hash.Pairs) != 0 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
	}
}

func TestParsingHashLiteralsWithExpressions(t *testing.T) {
	input := `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Value.(*ast.HashLiteral)

	if !ok {
		t.Fatalf("exp not ast.HashLiteral. got=%T", stmt.Value)
	}

	tests := map[string]func(ast.Expression){
		"one": func(e ast.Expression) {
			testInfixExpression(t, e, 0, "+", 1)
		},
		"two": func(e ast.Expression) {
			testInfixExpression(t, e, 10, "-", 8)
		},
		"three": func(e ast.Expression) {
			testInfixExpression(t, e, 15, "/", 5)
		},
	}

	for key, value := range hash.Pairs {
		literal, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key not ast.StringLiteral. got=%T", key)
		}

		testFunc, ok := tests[literal.Value]
		if !ok {
			t.Errorf("no test function for key %s found", literal.Value)
			continue
		}

		testFunc(value)
	}
}

func TestParsingHashLiteralsWithIntegerKey(t *testing.T) {
	input := `{1 : 2}`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	testNoErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Value.(*ast.HashLiteral)

	if !ok {
		t.Fatalf("exp not ast.HashLiteral. got=%T", stmt.Value)
	}

	tests := map[int64]func(ast.Expression){
		1: func(e ast.Expression) {
			testIntegerLiteral(t, e, int64(2))
		},
	}

	for key, value := range hash.Pairs {
		literal, ok := key.(*ast.IntegerLiteral)
		if !ok {
			t.Errorf("key not ast.IntegerLiteral. got=%T", key)
		}

		testFunc, ok := tests[literal.Value]
		if !ok {
			t.Errorf("no test function for key %d found", literal.Value)
			continue
		}

		testFunc(value)
	}
}
