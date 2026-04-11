package parser

import (
	"bufio"
	"fmt"
	"os"
)

func GenerateTokens(file *os.File) {
	// not actually doing anything yet, just prints the .stp file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
