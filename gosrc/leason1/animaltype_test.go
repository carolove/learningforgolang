package leason1_test

import (
	"fmt"
	"testing"

	"github.com/carolove/learningforgolang/gosrc/leason1"
)

func Test_AnimalType(t *testing.T) {
	var en leason1.AnimalEnumType = leason1.CAT
	if en == leason1.CAT {
		fmt.Println(en)
		en = leason1.DOG
	}
}
