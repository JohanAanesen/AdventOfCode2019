package main

import (
	"fmt"
	"os"
)

func main() {

	input := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,19,5,23,2,9,23,27,1,5,27,31,1,5,31,35,1,35,13,39,1,39,9,43,1,5,43,47,1,47,6,51,1,51,13,55,1,55,9,59,1,59,13,63,2,63,13,67,1,67,10,71,1,71,6,75,2,10,75,79,2,10,79,83,1,5,83,87,2,6,87,91,1,91,6,95,1,95,13,99,2,99,13,103,1,103,9,107,1,10,107,111,2,111,13,115,1,10,115,119,1,10,119,123,2,13,123,127,2,6,127,131,1,13,131,135,1,135,2,139,1,139,6,0,99,2,0,14,0}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {

			newInput := make([]int, len(input))
			copy(newInput, input)

			newInput[1] = i
			newInput[2] = j


			fmt.Println(newInput)

			//19690720
			if computer(newInput) == 19690720 {
				fmt.Println(i, j)
				os.Exit(0)
			}
		}
	}

	os.Exit(0) //quit



}

func computer(input []int) int {

	counter := 0;

	for input[counter] != 99 {
		if input[counter] == 1 {
			input[input[counter+3]] = input[input[counter+1]] + input[input[counter+2]]
		}else if input[counter] == 2 {
			input[input[counter+3]] = input[input[counter+1]] * input[input[counter+2]]
		}

		counter += 4
	}

	return input[0]
}