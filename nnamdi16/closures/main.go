package main
//go mode init closures
import (
	"fmt"
	"closures/fibonacci"
)


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
	n := 10
	for i := 0; i <= n; i++ {
		fmt.Printf("%d", fibonacci.Fibor(i))
	}
	fmt.Println()

	nextFibo := fibonacci.Fiboi()
	for i := 0; i <= n; i++ {
		fmt.Printf("%d", nextFibo())
	}
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
}