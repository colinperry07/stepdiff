package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Command flags
	inspect := flag.String("inspect", "", "-inspect=<path/to/file.stp>")
	flag.Parse()

	// Logic
	if *inspect != "" {
		file, err := os.Open(*inspect)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			line := scanner.Text()

			if !strings.HasPrefix(line, "#") {
				continue
			}

			lineId := line[strings.Index(line, "#"):strings.Index(line, "=")]
			lineType := line[strings.Index(line, "=")+1 : strings.Index(line, "(")]
			lineParams := line[strings.Index(line, lineType)+len(lineType) : strings.Index(line, ";")]

			strings.ReplaceAll(lineParams, "(", "")
			strings.ReplaceAll(lineParams, ")", "")

			params := strings.Split(lineParams, ",")

			var paramTypes []string

			switch lineType {
			case "CARTESIAN_POINT":
				paramTypes = append(paramTypes, "name", "coordinates")
			default:
				paramTypes = nil
				continue
			}

			fmt.Printf("%s %s\n", lineType, lineId)
			for idx, param := range params {
				if paramTypes != nil {
					fmt.Printf("%s: %s\n", paramTypes[idx], param)
				} else {
					continue
				}
			}

		}

		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading file:", err)
		}
	}

}
