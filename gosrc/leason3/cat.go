package leason3

import (
	"fmt"
)

type cat struct {
	name string
}

func NewCat(name string) Animal {
	ret := &cat{name}
	return ret
}

func (c *cat) Eat() {
	fmt.Println("fish")
}

func (c *cat) Name() string {
	return c.name
}
