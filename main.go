package main

import (
	"flag"
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
