// Package lexer implements functionality for converting fragments of the monkey programming language into tokens.
//
// For definitions of the tokens available in the monkey language see the [token] package.
package lexer

import "monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, l.ch)
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
