package leason1

type AnimalEnumType int
type Func func() (string, error)

const (
	_ AnimalEnumType = iota * 50
	CAT
	DOG
)

func NewFunc() {
	var fn Func = func() (string, error) {
		return "", nil
	}
	fn()
}
