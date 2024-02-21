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
