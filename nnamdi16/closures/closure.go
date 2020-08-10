package main

import "fmt"

// return a function which is returning int
// A closure is a function value (nextInt()) that references variables(i) from outside its body(intSeq()). But still remembers the value.
// The function may access and assign to the reference variables. The function is bound to the variables.
func intSeq() func() int {
	i := 0
	return func () int {
		i++
		return i
	}
}


func main()  {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
}