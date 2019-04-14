package main

import "fmt"

func main() {

	num := 17
	a := &num
	fmt.Println("Number is :", num)
	fmt.Println("Address is: ", a)
	fmt.Println("Number by deferencing pointer: ", *a)

	//Understanding concept of type of pointer
	val := hello()
	fmt.Println("Value returned is: ", *val)
	*val++
	hello()

	//Use slice of arrays when passing to modify array
}

func hello() *int {
	i := 10
	return &i
}
