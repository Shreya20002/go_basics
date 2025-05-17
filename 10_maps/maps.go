package main

import "fmt"

// maps -> hash, object, dict  in java, js, python
func main() {
	// creating maps

	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// getting an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["phone"]) // returns empty string if not found
    1:43:00
	

}
