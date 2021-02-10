package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string

	lname string
}

func main() {
	fmt.Print("Hello please write the name of a file:")

	var filename string

	fmt.Scan(&filename)

	names := make([]name, 0)

	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		temp := string(scanner.Text())

		separated := strings.Split(temp, " ")

		toAdd := name{separated[0], separated[1]}

		names = append(names, toAdd)
	}

	for _, value := range names {
		fmt.Printf("First Name: %s and Last name: %s\n", value.fname, value.lname)
	}
}
