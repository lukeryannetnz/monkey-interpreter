// Package lexer implements functionality for converting fragments of the monkey programming language into tokens.
//
// For definitions of the tokens available in the monkey language see the [token] package.
package lexer

import (
	"monkey-interpreter/token"
)

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
	l.eatWhitespace()

	switch l.ch {
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	case '=':
		peek := l.peekChar()
		if peek == '=' {
			tok.Type = token.EQ
			tok.Literal = string(l.ch) + string(peek)
			l.readChar()
		} else {
			tok = token.New(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '-':
		tok = token.New(token.MINUS, l.ch)
	case '!':
		peek := l.peekChar()
		if peek == '=' {
			tok.Type = token.NOT_EQ
			tok.Literal = string(l.ch) + string(peek)
			l.readChar()
		} else {
			tok = token.New(token.BANG, l.ch)
		}
	case '*':
		tok = token.New(token.ASTERISK, l.ch)
	case '/':
		tok = token.New(token.SLASH, l.ch)
	case '<':
		tok = token.New(token.LT, l.ch)
	case '>':
		tok = token.New(token.GT, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case '[':
		tok = token.New(token.LBRACKET, l.ch)
	case ']':
		tok = token.New(token.RBRACKET, l.ch)
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case ':':
		tok = token.New(token.COLON, l.ch)
	default:
		literal := l.readLiteral()
		tok = token.FindTokenType(literal)
	}

	l.readChar()
	return tok
}

func isWhitespace(ch byte) bool {
	// ascii values for tab, line feed, carriage return and space
	return ch == 9 || ch == 10 || ch == 13 || ch == 32
}

func isQuote(ch byte) bool {
	// ascii values for "
	return ch == 34
}

func (l *Lexer) eatWhitespace() {
	if isWhitespace(l.ch) {
		l.readChar()
		l.eatWhitespace()
	}
}

func (l *Lexer) readLiteral() string {
	literal := ""

	for l.ch != 0 {
		literal += string(l.ch)

		peek := l.peekChar()
		if isWhitespace(peek) || token.IsDelimiter(peek) {
			return literal
		}

		l.readChar()
	}

	return literal
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
