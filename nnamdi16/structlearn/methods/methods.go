package methods

import (
	"strconv"
	"structlearn/structsample"
)

//Defining methods on types
//A method is a function with a special receiver argument
//Example of a receiver - func (receiver) func_name(parameters) return_type{code}

//FirstPerson returns a Person struct
type FirstPerson struct {
	structsample.Person
}

//method (value receiver)

//Hello returns Hello
func (p FirstPerson) Hello() string  {
	return "Hello, I am " + p.FirstName + " " + p.LastName + ", " + strconv.Itoa(p.Age) + " years old"
}

// method (pointer receiver) - this modifies data

//HasBirthday returns age increment
func (p *FirstPerson) HasBirthday() {
	p.Age++
}