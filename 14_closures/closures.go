package main

import "fmt"

// closures are functions that can capture the environment in which they are created
// return func() int is created in the code block of count so it can use it inside
// it is not deleted after the function is executed ( ususally after fn is exec varisbles get deleted from callstack)

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}

}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
}
