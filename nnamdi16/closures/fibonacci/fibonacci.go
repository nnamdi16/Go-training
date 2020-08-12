package fibonacci


// recursive
func Fibor (n int) int {
	if n <= 1 {
		return n
	}
	return Fibor(n-1) + Fibor(n-2)
}

// Iterative
func Fiboi() func() int {
	x, y := 0, 1
	return func () int {
		r := x
		x , y = y, x + y

		return r
	}
}

