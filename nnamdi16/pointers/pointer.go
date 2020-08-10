package main


//Pointer holds the memory address of a value

import "fmt"

func main()  {
	//Assigning a value to  variable a
	a := 10

	//Assign the memory location of a where a is stored to b
	//The & operator generates a pointer to its operand.
	b := &a

	fmt.Println(a,b)

	//a is an int type while b is a pointer type(*int)
	fmt.Printf("%T %T\n", a,b)

	i, j := 42, 2701

	p := &i //points to i

	fmt.Println(*p) // read i through the pointer

	*p = 21  //set i through the pointer

	fmt.Println(i) //Check the new value of i

	p = &j //point to j
	*p = *p / 37  //divide j through the pointer

	fmt.Println(j) // see the new value of j




}