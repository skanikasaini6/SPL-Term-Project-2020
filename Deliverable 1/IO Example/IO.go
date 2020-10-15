package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte("This is IO in Go")
	fmt.Println("Writing to file")
	ioutil.WriteFile("file.data", data, 0777)
	fmt.Println("Write Succesful")
	fmt.Println("Reading from file")
	filedata, err := ioutil.ReadFile("file.data")
	if err != nil {
		fmt.Println("Catching and printing error")
		fmt.Println(err)
	}
	fmt.Println("Read Succesful")
	fmt.Print(string(filedata))

}
