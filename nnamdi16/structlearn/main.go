package main

import (
	"structlearn/structsample"
	"structlearn/methods"
	"structlearn/animal"
	"structlearn/ipaddress"
	"fmt"
)
// type Person struct {
// 	firstName string
// 	lastName  string
// 	city      string
// 	gender 	  string
// 	age  	  int
// }

func main()  {
	//Initialise person using struct 

	// p1 := Person{firstName:"Steven", lastName:"King", city:"Chicago",gender:"m",age:23}

	p1 := structsample.Person{FirstName:"Steven", LastName:"King", City:"Chicago",Gender:"m",Age:23}
	p2 := structsample.Person{"Neena", "Kocchar", "Boston", "f", 12}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.FirstName,p2.LastName)
	p2.Age++
	fmt.Println(p2)

	
	p3 := methods.FirstPerson{}

	p3.FirstName ="Stevens"
	p3.LastName ="Michael"
	p3.City="Lagos"
	p3.Gender="M"
	p3.Age =16

	p3.HasBirthday()

	fmt.Println(p3.Hello())
	fmt.Println(p3)

	a := animal.Animal{
		Name: "Gopher",
		Age: 3,
	}

	myStr := a.String()
	fmt.Println(myStr)


	host := map[string]ipaddress.IPAddr{
		"loopback" : {127, 0, 0, 1},
		"googleDNS" : {1,2,3,4},
	}

	for name,ip := range host {
		fmt.Printf("%v: %v\n", name, ip)
	}
}