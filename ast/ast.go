// Package ast provides an abstract syntax tree to represent monkey source code.
package ast

type Node interface {
	TokenLiteral() string
}

// A statement in Monkey does not produce a value
type Statement interface {
	Node
	statementNode()
}

// An expression in Monkey produces a value
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
