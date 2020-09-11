package main

import (
	"pointerreciever/point"
	"fmt"
)

func main()  {
	p := &point.Point{X:3,Y:4}
	fmt.Printf("before scaling : %+v, size: %v\n",p,p.Size())
	var sc float64 = 5
	p.Scale(sc)

	fmt.Printf("after scaling : %+v, size: %v\n",p,p.Size())
	
}