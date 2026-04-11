package parser

import (
	"bufio"
	"fmt"
	"os"
)

// token identifiers:
// keywords
// identifiers
// literals
// operators
// seperators

func GenerateTokens(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
