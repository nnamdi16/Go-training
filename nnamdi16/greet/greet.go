//Package greet derieved from  https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos
//https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go
package greet

import "fmt"

//Shark variable 
var Shark string = "Sammy"

//Octopus example for Defining a type 
type Octopus struct {
	Name string
	Color string
}

func (o Octopus) String() string {
	return fmt.Sprintf("The octopus's name is %q and is the color %s reserved.", o.Name, o.Color)
}

//Reset exported method
func (o *Octopus) Reset() {
	o.Name =""
	o.Color=""
}

//Hello prints Hello World
func Hello()  {
	fmt.Println("Hello, World!")
}