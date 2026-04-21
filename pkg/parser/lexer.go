package parser

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	SECTION_START = iota
	SECTION_END
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

var sectionStarters = map[string]bool{
	"ISO-10303-21": true,
	"HEADER":       true,
	"DATA":         true,
	"ANCHOR":       true,
	"REFERENCE":    true,
	"SIGNATURE":    true,
}

var sectionEnders = map[string]bool{
	"ENDSEC":           true,
	"END-ISO-10303-21": true,
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

func (l *Lexer) Next() (Token, error) {
	currentChar := l.peek(0)

	switch {
	case currentChar == 0:
		return Token{Type: EOF}, nil
	case currentChar == '#':
		return l.scanEntityRef()
	case unicode.IsUpper(currentChar):
		return l.scanKeyword()
	case currentChar == '\'':
		return l.scanString()
	case currentChar == '(':
		return Token{Type: LPAREN, Literal: "("}, nil
	default:
		return Token{}, nil
	}

}

func (l *Lexer) advance() rune {
	if l.pos >= len(l.input) {
		return rune(0)
	}

	_, runeSize := utf8.DecodeRune(l.input[l.pos:])
	l.pos += runeSize
	return l.peek(0)
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

	for unicode.IsUpper(l.peek(0)) || l.peek(0) == '_' || l.peek(0) == '-' {
		sb.WriteRune(l.advance())
	}

	word := sb.String()

	if sectionStarters[word] {
		return Token{Type: SECTION_START, Literal: word}, nil
	} else if sectionEnders[word] {
		return Token{Type: SECTION_END, Literal: word}, nil
	} else {
		return Token{Type: KEYWORD, Literal: word}, nil
	}

}

func (l *Lexer) scanString() (Token, error) {

	var sb strings.Builder

	l.advance() // consume opening quote

	for {
		char := l.advance()

		if char != '\'' {
			sb.WriteRune(char)
		} else if char == '\'' {
			if l.peek(0) == '\'' {
				sb.WriteRune(char)
				l.advance() // consume the second quote
			} else {
				break
			}
		}
	}

	return Token{Type: STRING, Literal: sb.String()}, nil

}
