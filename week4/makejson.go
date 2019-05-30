package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]string)

	var name, addr string

	fmt.Println("please enter a name")
	fmt.Scan(&name)

	fmt.Println("please enter an address")
	fmt.Scan(&addr)

	data["name"] = name
	data["address"] = addr

	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("can't marshal to json")
	}

	fmt.Printf("JSON object is:\n%s\n", string(json))
}
