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
		t.Errorf("letStmnt.Name.Vale not '%s'. got='%s", name, letStmnt.Name.Value)
		return false
	}

	if letStmnt.Name.TokenLiteral() != name {
		t.Errorf("letStmnt.Name.TokenLiteral() not '%s' got '%s'", name, letStmnt.Name.TokenLiteral())
		return false
	}

	return true
}
