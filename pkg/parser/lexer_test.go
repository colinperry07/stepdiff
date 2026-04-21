package parser_test

import (
	"reflect"
	"testing"

	"github.com/colinperry07/stepdiff/pkg/parser"
)

func TestTokenizeEntityRef(t *testing.T) {
	input := []byte("#123")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.ENTITY_REF, Literal: "#123"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenizeKeyword(t *testing.T) {
	input := []byte("TEST")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.KEYWORD, Literal: "TEST"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenizeSectionStart(t *testing.T) {
	input := []byte("HEADER")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.SECTION_START, Literal: "HEADER"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenizeSectionEnd(t *testing.T) {
	input := []byte("ENDSEC")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.SECTION_END, Literal: "ENDSEC"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenizeString(t *testing.T) {
	input := []byte("'hello'")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.STRING, Literal: "hello"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenizeStringWithEscapedQuote(t *testing.T) {
	input := []byte("'don''t'")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.STRING, Literal: "don't"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenizeLeftParen(t *testing.T) {
	input := []byte("(")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.LPAREN, Literal: "("},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenizeRightParen(t *testing.T) {
	input := []byte(")")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.RPAREN, Literal: ")"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}

func TestTokenizeMultiple(t *testing.T) {
	input := []byte("#123 TEST 'value'")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.ENTITY_REF, Literal: "#123"},
		{Type: parser.KEYWORD, Literal: "TEST"},
		{Type: parser.STRING, Literal: "value"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}
