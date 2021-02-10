package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Printf("Hi! Please type a name\n")

	fmt.Scan(&name)

	fmt.Printf("Please type an address too\n")

	fmt.Scan(&address)

	var toUse map[string]string = make(map[string]string)

	toUse["name"] = name

	toUse["address"] = address

	adr, _ := json.Marshal(toUse)

	fmt.Println(string(adr))
}
