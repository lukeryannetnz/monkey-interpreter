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
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// read two tokens so curToken and peekToken are populated
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		p.errors = append(p.errors, "failed to parse statement, found unexpected token of type: "+p.curToken.Literal)
		return nil
	}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		p.errors = append(p.errors, "unexpected token, expected IDENT")
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		p.errors = append(p.errors, "unexpected token, expected ASSIGN")
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !(p.curToken.Type == token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// assertion function to ensure the next token is of the expected type
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	} else {
		return false
	}
}
