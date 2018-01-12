package leason3_test

import (
	"fmt"
	"testing"

	"github.com/carolove/learningforgolang/gosrc/leason3"
)

func Test_Dog(t *testing.T) {
	ani := leason3.NewDog("beego")
	ani.Eat()
	fmt.Println(ani.Name())
}

func Test_Cat(t *testing.T) {
	ani := leason3.NewCat("lily")
	ani.Eat()
	fmt.Println(ani.Name())
}
