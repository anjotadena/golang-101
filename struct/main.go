package main

import (
	"fmt"
	"reflect"
)

// Struct is a collection of data fields with declared data types.

type Movie struct {
	Name   string
	Rating float32
}

type Hero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{time, "Klaxon"}
	return a
}

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	// Initialize struct movie
	m := Movie{
		Name:   "Citizen kane",
		Rating: 10,
	}

	// we can set value
	m.Name = "Hunter X Hunter"

	fmt.Println("Movie name: ", m.Name)
	fmt.Println("Movie rating: ", m.Rating)

	newMovie := Movie{}

	newMovie.Name = "Harry Potter"
	newMovie.Rating = 0.9232

	fmt.Println("NEW MOVIE")

	fmt.Println("Movie name: ", newMovie.Name)
	fmt.Println("Movie rating: ", newMovie.Rating)

	// Initializing a struct using new
	myMovie := new(Movie)
	myMovie.Name = "A"
	myMovie.Rating = 9

	fmt.Println("My Movie")
	fmt.Println("Movie name: ", myMovie.Name)
	fmt.Println("Movie rating: ", myMovie.Rating)

	oMovie := Movie{"Naruto", 10}

	fmt.Println(oMovie.Name)
	fmt.Println(oMovie.Rating)

	e := Hero{
		Name:    "Batman",
		Age:     32,
		Address: Address{1007, "Mountain Drive", "Gotham"},
	}

	// use dot notation to access nested struct
	fmt.Println(e.Address.City)

	// Default values for structs
	// Boolean false
	// Integer 0
	// Float 0.0
	// String ""
	// Pointer nil
	// Function nil
	// Interface nil
	// Slice nil
	// Channel nil
	// Maps nil

	// Setting default value
	// using constructor
	fmt.Println(NewAlarm("07:00"))

	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  false,
	}

	if a == b {
		fmt.Println("a and b are the same")
	}

	fmt.Println(a)
	fmt.Println(b)

	// Checking type
	fmt.Println("TypeOf a: ", reflect.TypeOf(a))
	fmt.Println("TypeOf b: ", reflect.TypeOf(b))

	c := a
	c.Ice = false
	fmt.Println("TypeOf a: ", a)
	fmt.Println("TypeOf b: ", b)
	fmt.Println("TypeOf c: ", c)

	// using pointer
	// to affect the changes
	d := &a

	d.Ice = false
	fmt.Println("TypeOf a: ", a)
	fmt.Println("TypeOf b: ", b)
	fmt.Println("TypeOf c: ", c)
	fmt.Println("TypeOf d: ", d)
}
