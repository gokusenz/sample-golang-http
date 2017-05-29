package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "Nobody expects the Spanish inquisition."
	re1, _ := regexp.Compile(`(expects (...) Spanish)`)
	result := re1.FindStringSubmatch(s)

	for k, v := range result {
		fmt.Printf("%d. %s\n", k, v)
	}
}
