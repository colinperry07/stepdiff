package parser

import (
	"fmt"
	"log"
	"os"
)

func ParseFile(filepath string) {

	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	l := Lexer{input: file, pos: 0}

	tokens, err := l.Tokenize()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tokens)

}
