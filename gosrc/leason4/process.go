package leason4

import "github.com/carolove/learningforgolang/gosrc/leason3"

type Proccess struct{}

func (p *Proccess) Proc(ani leason3.Animal) {
	ani.Eat()
}

func (p *Proccess) Proc1(fn func(context struct{})) {
	tx := struct{}{}
	fn(tx)
}
