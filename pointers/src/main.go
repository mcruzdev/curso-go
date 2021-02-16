package main

import "fmt"

func main() {
	a := 10

	var pointer *int = &a

	fmt.Println(*pointer)
}
