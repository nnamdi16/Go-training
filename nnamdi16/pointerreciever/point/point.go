package point

//Reasons to use a pointer receiver
//- So that the method can modfy the value that its receiver points to 
//- To avoid copying the value on each method call. It is more 
// efficient if the receiver is a large struct.

import (
	"math"
)
//Point type 
type Point struct {
	X,Y float64
}

func (p *Point) Scale(s float64) {
	p.X = p.X * s
	p.Y = p.Y * s
}

func (p *Point) Size() float64 {
	return math.Sqrt(p.X * p.X + p.Y * p.Y)
}

