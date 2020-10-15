package main

import "fmt"

var intArray [5]int          //This is an Array
var intSlice []int           //This is a Slice
var strintMap map[string]int //This is a Map
type marks struct {
	name  string
	id    int
	age   int
	marks int
} //This is a structure
func main() {

	intArray := [5]int{1, 2, 3, 4, 5} //Assigning values to Array
	intSlice = []int{100, 200, 300}   //Assigning values to Slice
	strintMap = map[string]int{
		"SpL":    100,
		"GoLang": 200,
	} //Assigning key and value to map
	student1 := marks{
		name:  "John Doe",
		id:    1337,
		age:   22,
		marks: 87,
	}

	fmt.Println("Printing the Integer Array: ", intArray)
	fmt.Println("Print the Integer Slice: ", intSlice)
	for x, y := range strintMap {
		fmt.Printf("%v has the value %d\n", x, y)
	}
	fmt.Println("Report for student: ", student1)
}
