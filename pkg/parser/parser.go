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

	l := New(file)

	tokens, err := l.Tokenize()
	if err != nil {
		log.Fatal(err)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

}
