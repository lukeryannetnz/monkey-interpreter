package parser

import (
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
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
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
return add(15);
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

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return'. got=%q", returnStmt.TokenLiteral())
		}
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