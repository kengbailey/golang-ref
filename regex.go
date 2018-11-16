package main

import (
	"fmt"
	"regexp"
)

func main() {

	// vars
	links := `
	https://gobyexample.com/string-formatting
	`

	// Compile regex formula first
	re := regexp.MustCompile(`https?://[0-9A-Za-z.-]+[0-9A-Za-z./#-]*/*`)

	// Look for matches
	linksParsed := re.FindAllString(links, -1)
	for i, v := range linksParsed {
		fmt.Printf("%v - %v", i, v)
	}

}
