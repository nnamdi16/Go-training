package main
import (
	"github.com/nnamdi16/Go-training/nnamdi16/greet"
	"fmt"
	"github.com/nnamdi16/Go-training/nnamdi16/morestrings"
	// "bufio"
	// "os"
	// "strconv"
)

func main()  {
	greet.Hello();
	fmt.Println(greet.Shark)

	oct:= greet.Octopus{
		Name:"Jesse",
		Color:"orange",
	}

	fmt.Println(oct.String())
	fmt.Println(oct.Name)
	oct.Reset()
	fmt.Println(oct.String())
	fmt.Println(morestrings.ReverseRunes("!oG , olleH"))

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type something: ")
	// scanner.Scan()
	// input := scanner.Text()
	// fmt.Printf("You typed: %q", input)
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type the year you were born:")
	// scanner.Scan()
	// input, _ := strconv.ParseIn t(scanner.Text(), 10, 64)
	// fmt.Printf("You will be %d years old at the end of 2020",2020-input)

	// var num1 float64 = 8
	// var num2 int = 4
	// var num3 int = 2
	// answer := num1/float64(num2)
	// answers := num2/num3
	// fmt.Printf("%g",answer)
	// fmt.Println("")
	// fmt.Printf("%d",answers)
	// fmt.Println("")
	// //Boolean

	// x := 5
	// y := 6.5
	// z := 9
	// val := x <= 5
	// val1 := float64(z) == y
	// fmt.Printf("%t", val)
	// fmt.Println("")
	// fmt.Printf("%t",val1)

	// name := "Tim"
	// fmt.Println("Before if")
	// if name != "tim" {
	// 	fmt.Println("Inside if")
	// }
	// fmt.Println("After if")
}

// package main

// import "fmt"

// func main()  {
// 	var number uint16 = 260
// 	fmt.Println("Hello World!",number)

// 	var numberType = 20000.00
// 	fmt.Printf("%T", numberType)

// 	//Another way to define a variable
// 	firstNumber := 2390
// 	fmt.Printf("%T" , firstNumber)

// 	//Default values
// 	//number uint64
// 	var secondFirst uint64
// 	var bl bool
// 	fmt.Println(secondFirst,bl)
// //Format %T - data type , %v - value  , %t - Boolean
// 	fmt.Printf("Hello %T %v %t", 10,10,true)
// 	fmt.Println("  ")

// 	//Format %b - Convert to base 2
// 	fmt.Printf("Number: %b", 3435)

// 	fmt.Println("  ")

// 	//Format %d - Convert to base 10
// 	fmt.Printf("Number: %d", 3435)

// 	fmt.Println("  ")

// 	//Format %x - Convert to base 16 (hexadecimal)
// 	fmt.Printf("Number: %x", 3435)

// 	fmt.Println("  ")
	
// 	//Format %e - exponential
// 	fmt.Printf("Number: %e", 3.8874737378383993)

// 	fmt.Println("  ")
	
// 	//Format %f - Convert to floating point
// 	fmt.Printf("Number: %f", 3.8874737378383993)

// 	fmt.Println("  ")
	
// 	//Format %g - Convert to floating point
// 	fmt.Printf("Number: %g", 3.8874737378383993)

// 	fmt.Println("  ")
// 	//Formatting string
// 	fmt.Printf("Number: %s", "Time")

// 	fmt.Println(" ")
// 	fmt.Printf("Number: %q", "Timely")

// 	fmt.Println(" ")
// 	//Precision
// 	fmt.Printf("Number: %-9q is cool", "tim")

// 	fmt.Println(" ")
// 	fmt.Printf("Number: %9q is cool", "tim")

// 	fmt.Println(" ") //Round to 2 decimal places
// 	fmt.Printf("Number: %.2f", 3.469829202)

// 	fmt.Println(" ")
// 	fmt.Printf("Number: %.f",89.98233)

// 	fmt.Println(" ")
// 	//Padding
// 	fmt.Printf("Number: %03d is cool", 45)

// 	fmt.Println(" ")
// 	//\n means new line
// 	var out string = fmt.Sprintf("Number: \n %07d is cool \n", 45)
// 	//\t- tab
// 	var inner string = fmt.Sprintf("Number: \t %07d is cool ", 45)
// 	fmt.Println(out)
// 	fmt.Println(inner)



// }
//To run a file using go run lesson1.go or go build lesson1.go and ./lesson1(build file)