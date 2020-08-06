//Range form of for loop iterates over a slice or map

package main

import "fmt"

type Cities struct {
	name string
	location [2] int
}

func rangeWithStruct()  {
	// Create empty slice of struct pointers.
	cities := []*Cities{}

	//create struct and appedn it to the slice.
	ct := new(Cities)
	ct.name = "London"
	ct.location[0] = 5
	ct.location[1] = 0
	cities = append(cities,ct)

	//create another struct
	ct = new(Cities)
	ct.name = "Sydney"
	ct.location[0] = 34
	ct.location[1] = 51
	cities = append(cities,ct)

	for i := range(cities) {
		c := cities[i]
		fmt.Println("City: ", *c)
	}
}

func rangeWithIndexesAndRunes()  {
	//range with string indexes and runes

	for i, c := range "ANVE" {
		fmt.Printf("%#U starts at byte position %d\n", c, i)
	}
}

func rangeWithMaps()  {
	//Range with a map
	moons := map[string]string{"Earth": "Moon", "Jupiter" :"Europa","Saturn" : "Titan"}
	
	for k, v := range moons {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func rangeWithMapsWithNumber()  {
	//range with a  map
	numbers := map[string]int{
		"Unos" : 1,
		"Dos" : 2,
		"Tres" : 3,
		"Cuatro" : 4,
		"Cinco" : 5,
	}

	for k, v := range numbers {
		fmt.Println(k,v)
	}
}

func rangeWithChannel()  {
	// We'll iterate over 5 values in the queue channel
	queue := make(chan string, 5)
	queue <- "Enceladus"
	queue <- "Titan"
	queue <- "Europa"
	queue <- "Ganemede"
	queue <- "IO"
	close(queue)

	//This `range` iterates over each element as it's
	// received from `queue`. Beacuse we `close`d the channel above 
	// channel above, the iteration terminates after  receiving the
	//receiving the 5 queues.
	for q := range queue {
		fmt.Println(q)
	}
}

func rangeWith()  {
	
}

func main()  {
	rangeWithStruct()
	rangeWithChannel()
	rangeWithMapsWithNumber()
	rangeWithMaps()
	rangeWithIndexesAndRunes()
	for _, i := range []int{1,2,3,4,5,6} {
		fmt.Println(i)
	}

	ids := []int {8,10,20,30,40,50,60}

	//Using index
	for i, id := range ids {
		fmt.Printf("%d -ID: %d\n", i,id)
	}

	//Sum of the ids
	sum := 0
	for _, id := range ids {
		sum +=id
	}
	fmt.Println("Sum", sum)
}
