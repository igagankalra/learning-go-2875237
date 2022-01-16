package main

import (
	"fmt"
)

func main() {
	anInt := 43.0
	p := &anInt
	fmt.Println("Value of p:", *p)

	value1 := 42.43
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)

	*pointer1 = anInt
	fmt.Println("new value of pointer 1:", *pointer1)
	fmt.Println("Memory address of pointer 1:", pointer1)
	fmt.Println("Memory address of anInt:", &anInt)
	fmt.Println("Memory address of anInt:", anInt)
}
