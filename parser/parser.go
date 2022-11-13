// Package parser contains functionality to syntactically analyse and convert tokens of the monkey programming language into the abstract syntax tree (AST) data structure.
//
// For more information on the AST, see the [ast] package.
//
// This parser uses a recursive descent approach, specifically it is a top down operator precedence or Pratt parser.
package parser

import (
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens so curToken and peekToken are populated
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
