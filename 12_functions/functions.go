package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// passing muliple arguments and returning multiple values with diff types
// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true
// }

// passing function as an argument
//func processIt(fn func(a int) int) {
//fn(1)
// }

// returning a function (like in js)
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	// kindaa like an arrow func in js
	// shorthand def of a function
	//  fn := func(a int) int {
	// 	return 2
	// }
	//  processIt(fn)

	fn1 := processIt()
	fmt.Println(fn1(5))
	fn1(6)

	// lang1, lang2, _ := getLanguages()
	// fmt.Println(lang1, lang2)
	// result := add(3, 5)
	// fmt.Println(result)
}
