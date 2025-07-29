package katas

import "fmt"

type day int

const (
	mon day = iota + 1 // iota can only be used for contants
	tue
	wed
	thu
	fri
	sat
	sun
)

func enums() {
	fmt.Println(mon)
	fmt.Println(tue)
	fmt.Println(sun)

	today := wed
	// var tomorrow int = thu // wont work because thu is constant
	var tomorrow day = thu // works

	switch today {
	case mon:
		fmt.Println("Monday")
	case tue:
		fmt.Println("Tuesday")
	case wed:
		fmt.Println("Wednesday")
	case thu | fri:
		fmt.Println("Alost there")
	case sat | sun:
		fmt.Println("Weekend :)")
	}

	fmt.Println(tomorrow.String())
}

// can also attach a string method to enums

func (d day) String() string {
	var r string
	switch d {
	case mon:
		r = "Monday"
	case tue:
		r = "Tuesday"
	case wed:
		r = "Wednesday"
	case thu:
		r = "Thursay"
	case fri:
		r = "Friday"
	case sat:
		r = "Saturday"
	case sun:
		r = "Sunday"
	}
	return r
}
