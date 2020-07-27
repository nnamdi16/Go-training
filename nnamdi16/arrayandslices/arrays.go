//A slice type is an abstraction built on top of Go's array type
//A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself. A slice is not an array but it describes a piece of an array
//A slice has both a length and a capacity. The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s):
package main

import "fmt"

func main()  {
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


}
