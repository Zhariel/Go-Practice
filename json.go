package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Print("Input your name: ")
	fmt.Scan(&name)
	fmt.Print("Input your address: ")
	fmt.Scan(&address)

	person := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
