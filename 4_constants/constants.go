package main

import "fmt"

const tot = 40

func main() {
	// constants
	// const is a keyword in golang to declare a constant
	// const pi = 3.14
	// const pi float64 = 3.14
	// const pi = 3.14 // this is a constant, it cannot be changed
	// fmt.Println(pi)
	// pi = 3.14159 // this will give an error, because pi is a constant
	// fmt.Println(pi)

	const (
		a = iota // iota is a predeclared identifier in golang, it is used to create enumerated constants
		b        // iota is incremented by 1 for each line
		c        // iota is incremented by 1 for each line
		d        // iota is incremented by 1 for each line
	)
	fmt.Println(a, b, c, d) // prints 0 1 2 3

	const (
		e = iota * 2 // e = 0 * 2 = 0
		f            // f = 1 * 2 = 2
		g            // g = 2 * 2 = 4
		h            // h = 3 * 2 = 6
	)
	fmt.Println(e, f, g, h) // prints 0 2 4 6
	const name = "golang"
	fmt.Println(name)
	// shorthand declaration cannot be done outside function, but normal and infer can be done
	//fmt.Println(agis)
	// error msg: syntax error: non-declaration statemet outside function body
	fmt.Println(tot)

	const (
		port = 8080
		host = "localhost"
	)
	fmt.Println(port, host)
}
