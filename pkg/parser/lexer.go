package parser

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	SECTION_START = iota // either HEADER or DATA
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

func New(input []byte) *Lexer {
	return &Lexer{input: input, pos: 0}
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
	currentChar := l.peek(0)

	switch {
	case currentChar == 0:
		return Token{Type: EOF}, nil
	case currentChar == '#':
		return l.scanEntityRef()
	case currentChar == '=':
		return l.scanKeyword()
	default:
		return Token{}, nil
	}

}

func (l *Lexer) peek(offset int) rune {
	if l.pos >= len(l.input) {
		return rune(0)
	}

	pos := l.pos
	for i := 0; i < offset; i++ {
		if pos >= len(l.input) {
			return rune(0)
		}
		_, runeSize := utf8.DecodeRune(l.input[pos:])
		pos += runeSize
	}

	if pos >= len(l.input) {
		return rune(0)
	}

	peekedRune, _ := utf8.DecodeRune(l.input[pos:])
	return peekedRune
}

func (l *Lexer) advance() rune {
	if l.pos >= len(l.input) {
		return rune(0)
	}

	peeked := l.peek(0)
	_, runeSize := utf8.DecodeRune(l.input[l.pos:])
	l.pos += runeSize
	return peeked
}

func (l *Lexer) scanEntityRef() (Token, error) {

	var sb strings.Builder

	// consumes the leading hashtag
	sb.WriteRune(l.advance())

	// consume digits until we hit a non-digit
	for unicode.IsDigit(l.peek(0)) {
		sb.WriteRune(l.advance())
	}

	return Token{Type: ENTITY_REF, Literal: sb.String()}, nil

}

func (l *Lexer) scanKeyword() (Token, error) {

	var sb strings.Builder

	// consume the leading equals sign and following uppercase letters
	sb.WriteRune(l.advance())

	for unicode.IsUpper(l.peek(0)) || l.peek(0) == '_' {
		sb.WriteRune(l.advance())
	}

	return Token{Type: KEYWORD, Literal: sb.String()}, nil
}
