package main

import (
	"flag"

	. "github.com/colinperry07/stepdiff/pkg/parser"
)

func main() {

	// Command flags
	inspect := flag.String("inspect", "", "-inspect=<path/to/file.stp>")
	flag.Parse()

	// Logic
	if *inspect != "" {
		ParseFile(*inspect)
	}

}
