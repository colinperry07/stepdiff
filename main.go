package main

import (
	"flag"

	. "github.com/colinperry07/stepdiff/pkg/parser"
)

func main() {

	// Command flags
	inspect := flag.String("inspect", "", "string=Path to CAD file to inspect")
	flag.Parse()

	// Logic
	if *inspect != "" {
		ParseFile(*inspect)
	} else {
		flag.Usage()
	}
}
