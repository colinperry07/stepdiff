package main

import "fmt"

// . "github.com/colinperry07/stepdiff/pkg/parser"

func main() {

	// // Command flags
	// inspect := flag.String("inspect", "", "-inspect=<path/to/file.stp>")
	// flag.Parse()

	// // Logic
	// if *inspect != "" {
	// 	ParseFile(*inspect)
	// }

	// testing for lexer

	line := "#14=CARTESIAN_POINT('',(-1.06945,-0.92765,0.184));"

	type TOKEN int

	const (
		POUND = iota
		NUM
		IDENT
		STRING
		EQUALS
		LPAREN
		RPAREN
		COMMA
		SEMICOLON
	)

	var _ = []string{
		POUND:     "#",
		NUM:       "NUM",
		STRING:    "STRING",
		IDENT:     "IDENT",
		EQUALS:    "=",
		LPAREN:    "(",
		RPAREN:    ")",
		COMMA:     ",",
		SEMICOLON: ";",
	}

	type Lexer struct {
		pos   int
		input string // holds the entire line
	}

	l := Lexer{pos: 0, input: line}

	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		fmt.Println(ch)
	}

}
