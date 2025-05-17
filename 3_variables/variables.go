package main

import "fmt"

func main() {
	var vari1 bool = true
	// if unused, compiler will give error, good for devs so that they don't unnecesarily declare redundant variables
	// can infer typ like python
	var vari = true
	fmt.Println(vari)
	fmt.Println(vari1)

	var age int = 20
	// golang automatically optimizes the memory allocation accoring to cpu architecture
	fmt.Println(age)

	// short hand declaration
	// := is a short hand declaration operator
	//name := "golang"
	//fmt.Println(name)

	// if we only declare ,, we need to mention the type, only assignment is done in shorthand without assign type
	var name string
	name = "golang"
	fmt.Println(name)

}
