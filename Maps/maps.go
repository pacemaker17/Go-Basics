package main

import "fmt"

func main() {

	studentsName := make(map[int]string)
	studentsName[1] = "Ashpreet"
	studentsName[2] = "Ashish"
	studentsName[3] = "Raghav"

	fmt.Println("Map is :", studentsName)

	// Maps are of reference types. Values get reflected if modified in any of them
	studentsCopy := studentsName
	studentsCopy[1] = "Deepu"

	fmt.Println("Students Name Map is :", studentsName)
	fmt.Println("Students Copy Map is :", studentsCopy)
}
