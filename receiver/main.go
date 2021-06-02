package main

import (
	"fmt"
	"os"
	"bufio"
)



// write code for receiver here
func main(){

	file, err := os.Open("generator_output.txt") // read our generator_output file in

	// handle event where the file cannot open correctly
	if err != nil {
		panic(err)
	}

	defer file.Close() // when the program is terminated, close our file

	scanner := bufio.NewScanner(file) // create our file handler object

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	} 

	// if there is an issue while reading the file
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}