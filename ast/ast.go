// Package ast provides an abstract syntax tree to represent monkey source code.
package ast

import "monkey-interpreter/token"

type Node interface {
	TokenLiteral() string
	String() string
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

// Program is the root of the AST which contains all other nodes
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

func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += s.String()
	}
	return out
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out string
	out += ls.TokenLiteral() + " "
	out += ls.Name.Value
	out += " = "

	if ls.Value != nil {
		out += ls.Value.String()
	}

	out += ";"
	return out
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out string
	out += rs.TokenLiteral() + " "

	if rs.Value != nil {
		out += rs.Value.String()
	} else {
		out += ";"
	}

	return out
}

type ExpressionStatement struct {
	Token token.Token
	Value Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	var out string
	out += es.TokenLiteral() + " "

	if es.Value != nil {
		out += es.Value.String()
	} else {
		out += ";"
	}

	return out
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}