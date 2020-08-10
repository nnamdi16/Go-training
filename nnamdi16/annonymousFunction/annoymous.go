package main

import "fmt"

//Regular Function
func sayHello(msg string)  {
	fmt.Println(msg)
}

// This is a regular function 
// name: regular_f_returning_annonymous_f
// return types: func(string) - annoymous function taking a string
func regularFunctionReturningAnonymousFunction() func(string) {
	// returns an anonymous function which is an inner functon
	return func (msg string)  {
		fmt.Println(msg)
	}
}

func main()  {
	sayHello("Hello from a regular function")

	//annonymous function
	func (msg string)  {
		fmt.Println(msg)
	}("Hello from an annonymous function")

	printFnc := regularFunctionReturningAnonymousFunction()

	printFnc("Hello from returned anonymous function")
}