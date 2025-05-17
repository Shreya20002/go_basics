package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop implementation using for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// infinite loop simulation
	//for{
	//println("1")
	//}
	// will give signal : interrupted due to infinite loop
	// type ctrl + c to stop the execution

	// classic for loop
	for i := 0; i < 4; i++ {
		//break
		// will give error
		if i == 2 {
			continue
			// will skip the iteration
		}
		fmt.Println(i)

	}

	// another for loop like in python
	for i := range 3 {
		fmt.Println(i)
		// will print uptill 2
	}

}
