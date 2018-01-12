package leason2

import (
	"fmt"
)

var Animal *animal

func init() {
	Animal = createAnimal("cat")
	fmt.Println("init Animal at Init")
}

type animal struct {
	name string
}

func createAnimal(name string) *animal {
	ret := &animal{name: name}
	return ret
}

func (an *animal) PrintName() {
	fmt.Println(an.name)
}
