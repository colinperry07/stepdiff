package parser

import (
	"fmt"
	"log"
	"os"
)

func ParseFile(filepath string) {

	// os.ReadFile() returns a byte slice as opposed to os.Open()
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	l := Lexer{input: file, pos: 0}

	tokens, err := l.Tokenize()
	if err != nil {
		log.Fatal(err)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

}
