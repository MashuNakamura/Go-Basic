package main

import (
	"fmt"
	"reflect"
)

func main() {
	var complex = 1 + 1
	var tipe = reflect.TypeOf(complex)

	fmt.Println("This is a sample Go program.")
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga = ", 3)
	fmt.Println(complex)
	fmt.Println(tipe)
}
