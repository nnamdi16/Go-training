package cities

import "fmt"

//Cities struct 
type Cities struct {
	Name string
	Location [2]int
}

//CreateCities method 
func CreateCities() {
	//create empty slice of struct pointers
	cities := []*Cities{}

	//Create struct and append it to the slice
	ct := new(Cities)
	ct.Name="London"
	ct.Location[0] = 5
	ct.Location[1] = 0
	cities = append(cities,ct)
	
	//create another struct
	ct = new(Cities)
	ct.Name = "Sydney"
	ct.Location[0] = 23
	ct.Location[1] = 51
	cities = append(cities,ct)

	for i := range (cities) {
		c := cities[i]
		fmt.Println("City: ", *c)
	}
}