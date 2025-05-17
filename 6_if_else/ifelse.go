package main

import "fmt"

func main() {
	/*age := 16
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
	*/
	//age := 16
	//if age >= 18 {
	//fmt.Println("You are an adult")
	//} else if age >= 13 {
	//fmt.Println("You are a teenager")
	//} else {
	//fmt.Println("You are a child")
	//}
	// if you're not using imported package anywhere, it automatically gets deleted , due to the extention we downloaded -> where the compiler does this task
	var role = "admin"
	var hasPermissions = true
	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}
	// and condition is written using && operator
	// or condition is written using || operator
	// not condition is written using ! operator

	// you can declare age := 16  (shorthand declaration) inside if else block and use it for if else conditions there
	if age := 16; age >= 18 {
		fmt.Println("You are an adult", age)
	} else if age >= 13 {
		fmt.Println("You are a teenager", age)
	} else {
		fmt.Println("You are a child", age)
	}

	// go does not have ternary operator like c, c++, java, python
	// A ternary operator is a shorthand way of writing an if-else statement in programming. It's called "ternary" because it involves three operands: a condition, a value if the condition is true, and a value if the condition is false.

	/*The syntax looks like this:


	      condition ? value_if_true : value_if_false;
		in c lang
	*/
}
