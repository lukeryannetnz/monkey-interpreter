// Package token provides definitions of tokens available in the monkey programming language.
package token

import (
	"strconv"
	"strings"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func FindTokenType(literal string) Token {
	if _, err := strconv.Atoi(literal); err == nil {
		return Token{Type: INT, Literal: literal}
	}

	if strings.EqualFold(LET, literal) {
		return Token{Type: LET, Literal: literal}
	} else {
		return Token{Type: IDENT, Literal: literal}
	}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // An identifier e.g. add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
