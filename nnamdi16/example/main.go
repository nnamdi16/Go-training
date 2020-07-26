package main
import (
	"github.com/nnamdi16/Go-training/nnamdi16/greet"
	"fmt"
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
}