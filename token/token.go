// Package token provides definitions of tokens available in the monkey programming language.
package token

import (
	"strconv"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

var delimiters = map[string]TokenType{
	",": COMMA,
	";": SEMICOLON,
	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
	"[": LBRACKET,
	"]": RBRACKET,
}

func New(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func FindTokenType(literal string) Token {
	if _, err := strconv.Atoi(literal); err == nil {
		return Token{Type: INT, Literal: literal}
	}

	if tok, ok := keywords[literal]; ok {
		return Token{Type: tok, Literal: literal}
	}

	return Token{Type: IDENT, Literal: literal}
}

func IsDelimiter(literal byte) bool {
	_, ok := delimiters[string(literal)]

	return ok
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT  = "IDENT" // An identifier e.g. add, foobar, x, y, ...
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)
