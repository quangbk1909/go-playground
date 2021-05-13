package pkga

import "fmt"

type A interface {
	testA()
}

type ExportedA []unexportA

type unexportA struct {

}

func  NewUnexportA() unexportA {
	return unexportA{}
}


func (a unexportA) testA() {
	fmt.Println("aaa")
}
