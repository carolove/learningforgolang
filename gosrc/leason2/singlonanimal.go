package leason2

import (
	"fmt"
	"sync"
)

// single with init
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

// single with once
var ani *animal
var once sync.Once

func GetAnimalInstance(name string) *animal {

	once.Do(func() { ani = &animal{name} })
	return ani
}
