package parser

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func ParseFile(filepath string) {

	redundantLines := []string{"ISO-10303-21;", "END-ISO-10303-21;"}

	dirtyFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var cleanFile []byte

	for i := 0; i < len(redundantLines); i++ {
		editedFile := bytes.ReplaceAll(dirtyFile, []byte(redundantLines[i]), []byte(""))
		cleanFile = editedFile
	}

	l := New(cleanFile)

	tokens, err := l.Tokenize()
	if err != nil {
		log.Fatal(err)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

}
