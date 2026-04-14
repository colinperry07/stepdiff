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
	input := []byte("=TEST")
	got, err := parser.New(input).Tokenize()
	if err != nil {
		t.Fatalf("Tokenize() returned error: %v", err)
	}

	want := []parser.Token{
		{Type: parser.KEYWORD, Literal: "=TEST"},
		{Type: parser.EOF, Literal: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Tokenize() = %#v, want %#v", got, want)
	}
}
