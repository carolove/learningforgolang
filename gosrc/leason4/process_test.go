package leason4_test

import (
	"fmt"
	"testing"

	"github.com/carolove/learningforgolang/gosrc/leason3"
	"github.com/carolove/learningforgolang/gosrc/leason4"
)

func Test_Proc(t *testing.T) {
	proc := &leason4.Proccess{}
	ani := leason3.NewDog("beego")
	proc.Proc(ani)

	ani = leason3.NewCat("lily")
	proc.Proc(ani)
}

func Test_Proc1(t *testing.T) {
	proc := &leason4.Proccess{}
	proc.Proc1(func(context struct{}) {
		fmt.Println("i eat bones")
	})
	proc.Proc1(func(context struct{}) {
		fmt.Println("i eat fish")
	})
}
