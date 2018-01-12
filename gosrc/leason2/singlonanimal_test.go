package leason2_test

import (
	"fmt"
	"testing"

	"github.com/carolove/learningforgolang/gosrc/leason2"
)

func Test_Singlon(t *testing.T) {
	leason2.Animal.PrintName()
	fmt.Println(&leason2.Animal)
	fmt.Println(&leason2.Animal)
}

func Test_SinglonOnce(t *testing.T) {
	animal1 := leason2.GetAnimalInstance("cat")
	animal1.PrintName()
	animal2 := leason2.GetAnimalInstance("dog")
	animal2.PrintName()
}
