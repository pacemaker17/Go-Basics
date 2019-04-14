package main

import "fmt"

func main() {

	s := "Ashpreet"
	//afterRuned := []rune(s)
	//fmt.Print(afterRuned)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	for _, rune := range s {
		fmt.Printf("%c ", rune)
	}
}
