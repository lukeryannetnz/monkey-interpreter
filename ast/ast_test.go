package ast

import (
	"monkey-interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foo"},
					Value: "foo",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "bar"},
					Value: "bar",
				},
			},
		},
	}

	if program.String() != "let foo = bar;" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}

func TestIfExpressionString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ExpressionStatement{
				Token: token.Token{Type: token.IF, Literal: "IF"},
				Value: &IfExpression{
					Token: token.Token{Type: token.IF, Literal: "if"},
					Condition: &PrefixExpression{
						Token:    token.Token{Type: token.BANG, Literal: "!"},
						Operator: "!",
						Right: &Boolean{
							Token: token.Token{Type: token.FALSE, Literal: "false"},
							Value: true},
					},
					Consequence: &BlockStatement{
						Statements: []Statement{
							&LetStatement{
								Token: token.Token{Type: token.LET, Literal: "let"},
								Name: &Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "foo"},
									Value: "foo",
								},
								Value: &Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "bar"},
									Value: "bar",
								},
							},
						},
					},
				},
			},
		},
	}

	if program.String() != "if(!false) let foo = bar;" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}

func TestTokenLiteral(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foo"},
					Value: "foo",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "bar"},
					Value: "bar",
				},
			},
		},
	}

	tl := program.TokenLiteral()

	if tl != "let" {
		t.Fatalf("token literal not correct. expected=let, got=%s", tl)
	}
}

func TestCallExpressionString(t *testing.T) {
	ce := &CallExpression{
		Token: token.Token{Type: token.LPAREN, Literal: "("},
		Function: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "foo"},
			Value: "foo",
		},
		Arguments: []Expression{
			&StringLiteral{
				Token: token.Token{Type: token.IDENT, Literal: "bar"},
				Value: "bar",
			},
		},
	}

	if ce.String() != "foo(bar)" {
		t.Fatalf("string not correct. expected=foo(bar), got=%s", ce.String())
	}
}

func TestFunctionLiteralString(t *testing.T) {
	ce := &FunctionLiteral{
		Token: token.Token{Type: token.FUNCTION, Literal: "fn"},
		Parameters: []*Identifier{
			{
				Token: token.Token{Type: token.IDENT, Literal: "foo"},
				Value: "foo",
			},
		},
		Body: &BlockStatement{
			Token: token.Token{Type: token.LBRACE, Literal: "{"},
			Statements: []Statement{
				&ExpressionStatement{
					Token: token.Token{Type: token.IDENT, Literal: "bar"},
					Value: &StringLiteral{
						Token: token.Token{Type: token.IDENT, Literal: "bar"},
						Value: "bar",
					},
				},
			},
		},
	}

	if ce.String() != "fn(foo){ bar }" {
		t.Fatalf("string not correct. expected=fn(foo){ bar }, got=%s", ce.String())
	}
}

func TestArrayLiteralString(t *testing.T) {
	sut := &ArrayLiteral{
		Token: token.Token{},
		Elements: []Expression{
			&StringLiteral{
				Token: token.Token{Type: token.IDENT, Literal: "foo"},
				Value: "foo",
			},
			&StringLiteral{
				Token: token.Token{Type: token.IDENT, Literal: "bar"},
				Value: "bar",
			},
		},
	}

	result := sut.String()

	if result != "[foo, bar]" {
		t.Fatalf("string not as expected. expected:[foo, bar] got:%s", result)
	}
}
