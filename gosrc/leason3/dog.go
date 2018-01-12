package leason3

import (
	"fmt"
)

func NewDog(name string) Animal {
	ret := &dog{name}
	return ret
}

type dog struct {
	name string
}

func (d *dog) Eat() {
	fmt.Println("bone")
}

func (d *dog) Name() string {
	return d.name
}
