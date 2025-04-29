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
	without_var := "Without var"
	test := "Hello World!"
	var (
		first_str = "First String"
		second_str = "Second String"
	)

	// Example Const : var that can't be changed or modify
	const const_acumalaka = "You cannot change this"

	// Example of a number variable
	fmt.Println("This is a sample Go program.")
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga = ", 3)
	fmt.Println(first_str)
	fmt.Println(second_str)

	fmt.Println()
	// Example calling a string variable
	fmt.Println(complex)
	fmt.Println(tipe)
	acumalaka = "Named Var"
	fmt.Println(acumalaka)
	fmt.Println(without_var)
	fmt.Println(test)

	fmt.Println()

	fmt.Println(const_acumalaka)

	fmt.Println()

	// Example of a boolean variable
	fmt.Println("True", true)
	fmt.Println("False", false)

	// Get the length of a string
	fmt.Println(len("Hello World!"))

	// Get the index of the first character in a string
	fmt.Println("Hello World"[0])
}
