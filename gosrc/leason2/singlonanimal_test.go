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
