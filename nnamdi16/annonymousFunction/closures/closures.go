package closures

// GO supports annonymous functions which can form closure - We can define a function without having to name it.
// Go functions may be closures.


// IntSeq return a function which is returning int
func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


