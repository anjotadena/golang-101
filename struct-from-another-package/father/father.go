package father

import "fmt"

type Father struct {
	Name string
}

func init() {
	fmt.Println("Father package initialized.")
}

func (f Father) Data(name string) string {
	f.Name = name

	return "Father: " + f.Name
}
