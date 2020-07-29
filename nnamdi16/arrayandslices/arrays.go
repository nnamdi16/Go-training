//A slice type is an abstraction built on top of Go's array type
//A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself. A slice is not an array but it describes a piece of an array
//A slice has both a length and a capacity. The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s):
package main

import "fmt"

func leftRotation(a []int, size int, rotation int) []int {
	var newArray []int

	for i := 0; i < rotation; i++ {
		newArray = a[1:size]
		fmt.Println(newArray)
		newArray = append(newArray,a[0])
		fmt.Println(newArray)
		a = newArray
	}
	return a
}

func main()  {
	a := [] int {1,2,3,4,5}
	// fmt.Println(len(a))
	fmt.Printf("input : %+v\n", a)
	rotation := 4
	fmt.Printf("output: %+v\n", leftRotation(a, len(a),rotation))
	var fruits [2]string
	fruits[0] = "pomegranate"
	fruits[1] = "rambutan"
	continents := [2]string{"Asia", "Africa"}
	ids := []int{1,3,5,7,9}
	alphabet := []string{"a","b","c","d","e"}
	s:= make([]int, 5,10)
	fmt.Printf("len=%d cap=%d %v\n", len(s),cap(s),s)
	fmt.Println(alphabet);
	fmt.Println(continents);
	fmt.Println(fruits);
	fmt.Println(ids);

	fruitsSlice := []string{"pomegranate", "rambutan", "mangosteen","jackfruit"}

	fmt.Println(fruitsSlice)
	fmt.Println(len(fruitsSlice))
	fmt.Println(fruitsSlice[1:3])
	fmt.Println(fruitsSlice[1:4])
	fmt.Println(fruitsSlice[3:])
	fmt.Println(fruitsSlice[:])

	//Byte slice
	sByte := []byte("abc")

	//Print initial byte
	fmt.Println(sByte)

	//append  a byte
	sByte = append(sByte, byte('d'))

	//Print string representation of the byte
	fmt.Println(string(sByte))

	//length of the byte slice
	fmt.Println(len(sByte))

	//1st byte
	fmt.Println(string(sByte[0]))
//Differences between array and slices
//Array is fixed while slices is dynamic
//The number of elments in a slice can grow dynamically
//Slice uses array behind the curtain but slice by itself cannot store any data
//Slice plays with a portion of the underlying array
// The length of a slice is the length of the segment of the array that the slice contains.
// The capacity is the maximum size up to which the segment can grow

var array = [6]int{10,20,30,40,50,60}
var arrayslice = array[1:4]
fmt.Printf("slice: cap %v, len %v, %v\n",cap(arrayslice),len(arrayslice), arrayslice)
fmt.Printf("array: cap %v, len %v, %v\n", cap(array), len(array), array)

//make function is used to create slices
//The make() function takes a type, a length, and an optional capacity.
//It allocates an underlying array with size equal to the given capacity, and returns a slice that refers to that array:
freeSlice := make([]int,4,10)
fmt.Printf("len = %v, cap = %v, slice = %v\n", len(freeSlice),cap(freeSlice),freeSlice)

}
