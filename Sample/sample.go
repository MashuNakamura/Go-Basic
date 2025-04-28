package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Example of a string variable
	var complex = 1 + 1
	var tipe = reflect.TypeOf(complex)
	var acumalaka string
	test := "Hello World!"

	// Example of a number variable
	fmt.Println("This is a sample Go program.")
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga = ", 3)

	fmt.Println()
	// Example calling a string variable
	fmt.Println(complex)
	fmt.Println(tipe)
	fmt.Println("Penamaan Var", acumalaka)
	fmt.Println(test)

	fmt.Println()

	// Example of a boolean variable
	fmt.Println("True", true)
	fmt.Println("False", false)

	// Get the length of a string
	fmt.Println(len("Hello World!"))

	// Get the index of the first character in a string
	fmt.Println("Hello World"[0])
}
