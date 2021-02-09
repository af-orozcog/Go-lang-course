package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string

	fmt.Scanf("%q", &word)

	word = strings.ToLower(word)

	if strings.HasPrefix(word, "i") && strings.HasSuffix(word, "n") && strings.Contains(word, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
