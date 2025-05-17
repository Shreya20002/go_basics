package main

import (
	"fmt"
	"time"
)

func main() {
	// switch case
	// switch is a multi branch statement
	// switch is like if else if else
	// switch is like a dictionary in python
	// switch is like a match case in python
	// switch is like a match case in c++
	// switch is like a match case in java
	// switch is like a match case in javascript
	// switch is like a match case in c#
	// switch is like a match case in ruby
	// switch is like a match case in php
	// switch is like a match case in go
	var i int = 1
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is not 1, 2 or 3")
	}
	// unlike c++ we don't have to add break statements after every case in go , it handles that automatically , even providing default is not compulsory

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	// today is friday , hence it will print "It's a weekday"

	// type switch
	// function behaves like a first class citizen in go, just like in python
	// function can be passed as an argument to another function
	// i interface{} meaning any type of value can be taken up by i
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an int")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a bool")
		default:
			fmt.Println("I don't know", t)
		}
		// if you dont want to print or use t just do switch i.(type), it will run without t now
	}
	whoAmI("golang")
}
