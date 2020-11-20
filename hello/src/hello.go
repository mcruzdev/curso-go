package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Matheus"
	var version float32 = 1.1

	nameInf := "Cruz"
	versionInf := 1.3

	fmt.Println("Hello, Sir", name)
	fmt.Println("Application's version: ", version)

	fmt.Println("Variable type is: ", reflect.TypeOf(name))
	fmt.Println("Variable type is: ", reflect.TypeOf(version))
	fmt.Println("Variable type is: ", reflect.TypeOf(nameInf))
	fmt.Println("Variable type is: ", reflect.TypeOf(versionInf))
}
