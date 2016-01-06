package main

import "fmt"

type myStruct struct {
	Field int
}

func main() {

	var b myStruct = myStruct{Field: 1}

	c := struct {
		Field string
	}{
		Field: "Hello",
	}

	var a float64 = 1.0 - 0.000000000000000001

	fmt.Printf("a: %.21f\nb: %#v\nc: %#v", a, b, c)
}
