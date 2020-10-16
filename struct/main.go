package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// Struct is a collection of data fields with declared data types.

type Movie struct {
	Name   string
	Rating float64
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

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)

	return m.Name + ", " + r
}

type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius

	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

type Triangle struct {
	Base   float64
	Height float64
}

// Passing a Pointer Reference to a Method
func (t *Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

// Passing a Value Reference to a Method
func (t Triangle) ChangeBase(f float64) {
	// Any changes on this will not be affected to the original value
	// this function operates a copy of a variable
	t.Base = f
	return
}

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	}
	return nil
}

func Boot(r Robot) error {
	return r.PowerOn()
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

	// Declaring methods

	fmt.Println("Movie Summary: ", m.summary())

	sphere := Sphere{
		Radius: 5,
	}

	fmt.Println("Sphere surface area: ", sphere.SurfaceArea())
	fmt.Println("Sphere volume: ", sphere.Volume())

	t := Triangle{Base: 3, Height: 1}
	fmt.Println("Triangle Area: ", t.Area())

	t.ChangeBase(4) // operates a copy value
	fmt.Println("Triangle Change Base: ", t)

	// NOTES:
	// If you want to modify or mutate the original initialization of a struct use pointer else use value ref

	// USING INTERFACES
	// - specifies a method set
	// - as Blueprint of your struct
	// - describe all the methods of a method to set and provides the function signatures for each method.

	fmt.Println("*************ROBOTS*****************")

	r1 := T850{
		Name: "The terminator",
	}

	r2 := R2D2{
		Broken: true,
	}

	err := Boot(&r2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&r1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}
}
