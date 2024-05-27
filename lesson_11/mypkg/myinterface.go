package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("methodWithoutParameters called")
}

func (m MyType) MethodWithParameter(param float64) {
	fmt.Println("methodWithParameter called", param)
}

func (m MyType) MethodWithReturnValue() string {
	return "methodWithReturnValue called"
}

func (m MyType) MethodNotInterface() {
	fmt.Println("MethodNotInterface called")
}
