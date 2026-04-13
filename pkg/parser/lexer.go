package parser

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	SECTION_START = iota
	ENTITY_REF
	KEYWORD
	STRING
	INTEGER
	REAL
	ENUM
	BINARY
	LPAREN
	RPAREN
	COMMA
	SEMICOLON
	EQUALS
	DOLLAR
	ASTERISK
	COMMENT
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input []byte
	pos   int
}

func (l *Lexer) Tokenize() ([]Token, error) {

	var tokens []Token

	for {

		token, err := l.Next()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)

		if token.Type == EOF {
			break
		}
	}

	return tokens, nil

}

func (l *Lexer) Next() (Token, error) {
	// current rune access is l.peek(0)

	currentChar := l.peek(0)

	switch {
	case currentChar == '#':
		return l.scanEntityRef()
	default: // placeholder
		return Token{}, nil
	}

}

// need New() func

func (l *Lexer) peek(offset int) rune {

	var runeSize int

	pos := l.pos
	for i := 0; i < offset; i++ {
		_, runeSize = utf8.DecodeRune(l.input[pos:])
		pos += runeSize
	}

	peekedRune, _ := utf8.DecodeRune(l.input[pos:])

	return peekedRune
}

func (l *Lexer) advance() rune {
	peeked := l.peek(0)
	l.pos++
	return peeked
}

func (l *Lexer) scanEntityRef() (Token, error) {

	var sb strings.Builder

	// consumes the leading hashtag(pound)
	sb.WriteRune(l.advance())

	for unicode.IsDigit(l.peek(0)) {
		sb.WriteRune(l.advance())
	}

	return Token{Type: ENTITY_REF, Literal: sb.String()}, nil

}
