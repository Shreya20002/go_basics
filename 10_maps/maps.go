package main

import "fmt"

// maps -> hash, object, dict  in java, js, python
func main() {
	// creating maps

	//m := make(map[string]string)

	// setting an element
	//m["name"] = "golang"
	//m["area"] = "backend"

	// getting an element
	//fmt.Println(m["name"], m["area"])
	//fmt.Println(m["phone"]) // returns empty string if not found , 0 if it is an int not found
	//1:43:00
	// shift + alt + drag mouse to select multiple lines and change smth simultaneously
	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 50
	fmt.Println(m["phone"])
	fmt.Println(len(m))

	delete(m, "price")
	clear(m)

	fmt.Println(m)
	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phones": 3}

	// v, ok := m["phones"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// m1 := map[string]int{"price": 40, "phones": 3}
	// m2 := map[string]int{"price": 40, "phones": 8}
	// fmt.Println(maps.Equal(m1, m2))
	// maps package from std lib to get equal function to comapre as map is an object
	//

}
