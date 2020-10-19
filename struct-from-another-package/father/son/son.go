package son

import "fmt"

type Son struct {
	Name string
}

func init() {
	fmt.Println("Son package initialized")
}

func (s Son) Data(name string) string {
	s.Name = name

	return "Son: " + s.Name
}
