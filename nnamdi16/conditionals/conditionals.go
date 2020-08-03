package main

import "fmt"
func switchCase()  {
	galaxy := "M87"
	switch galaxy {
	case "Milky Way":
		fmt.Printf("Galaxy name is 'Milky Way' \n")
	case "Andromeda":
		fmt.Printf("Galaxy name is 'Andromedia'\n")
	case "M87":
		fmt.Printf("Galaxy name is 'M87'\n")
	}
}

func main()  {
	a, b := 1,10

	if a < b {
		fmt.Printf("%d is less than %d\n", a,b)
	} else if a == b {
		fmt.Printf("%d is equal to %d\n", a, b)
	} else {
		fmt.Printf("%d is greater than %d\n", a, b)
	}

	switchCase()
	name := "tim"
	fmt.Println("Before if")
	if name != "tim" || name == "tim" {
		fmt.Println("Inside if")
	}
	fmt.Println("After it")
}