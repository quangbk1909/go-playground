package pkgb

import "fmt"

type B interface {
	TestB()
}

type unexportB struct {

}

func  NewUnexportB() *unexportB {
	return &unexportB{}
}


func (a *unexportB) TestB() {
	fmt.Println("aaa")
}
