package parser

import (
	"log"
	"os"
)

func ParseFile(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	GenerateTokens(file)
}
