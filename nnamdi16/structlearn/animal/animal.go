package animal

import (
	"fmt"
)

// Animal has a Name and an Age to represent an animal
type Animal struct {
	Name string
	Age uint
}

//String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}