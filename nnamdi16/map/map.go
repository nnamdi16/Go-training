package main

//Hash table offers lookups, adds, deletes. Go provides a built-in map type that implements a hash table.


import "fmt"

func isBalanced(p string) bool {
	pairs := map[rune]rune{'(' : ')', '{':'}', '[':']'} // rune - alias for int32

	fmt.Println(pairs)
	stack := []rune{}

	for _, c := range p {
		if c =='(' || c == '{' || c == '[' {
			stack = append(stack, c)
			fmt.Println("We wait")
			fmt.Println(stack)
		} else {
			if len(stack) < 1 {
				return false
			}

			if c == pairs[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return true
}

func main()  {
	
	//Define map
	moons := make(map[string]string)

	//Assign
	moons["Earth"] = "Moon"
	moons["Jupiter"] = "Europe"
	moons["Saturn"] = "Titan"

	fmt.Println(moons)

	//Delete 
	delete(moons, "Saturn")
	fmt.Println(moons)


	planet := map[string]string{
		"Earth":"Moon",
		"Jupiter":"Europe",
		"Saturn":"Titan",
	}

	fmt.Println(planet)

	parentheses :=
		[]string{")(", "[(])", "(()){}[]", "())[]{}", "()[]{}(([])){[()][]}"}

	for _, p := range parentheses {
		fmt.Printf("%s: %t\n", p, isBalanced(p))
	}

}