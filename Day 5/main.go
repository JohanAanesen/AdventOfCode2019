package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := []int{3,8,1001,8,10,8,105,1,0,0,21,34,59,76,101,114,195,276,357,438,99999,3,9,1001,9,4,9,1002,9,4,9,4,9,99,3,9,102,4,9,9,101,2,9,9,102,4,9,9,1001,9,3,9,102,2,9,9,4,9,99,3,9,101,4,9,9,102,5,9,9,101,5,9,9,4,9,99,3,9,102,2,9,9,1001,9,4,9,102,4,9,9,1001,9,4,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,99}

	fmt.Println(computer(input))


	os.Exit(0) //quit

}

func computer(input []int) int {

	counter := 0

	for true {

		paddedIntCode := fmt.Sprintf("%05d",input[counter])
		param1PositionMode := paddedIntCode[2:3] == "0"
		param2PositionMode := paddedIntCode[1:2] == "0"

		opCode, _ := strconv.Atoi(paddedIntCode[3:5])

		if opCode == 99 {
			return 69
		}

		if opCode == 1 || opCode == 2 {

			param1 := input[counter+1]
			if param1PositionMode {
				param1 = input[input[counter+1]]
			}

			param2 := input[counter+2]
			if param2PositionMode {
				param2 = input[input[counter+2]]
			}

			if opCode == 1 {
				input[input[counter+3]] = param1 + param2
			}else if opCode == 2 {
				input[input[counter+3]] = param1 * param2
			}

			counter += 4

		}else if opCode == 3 {
		 // input
			input[input[counter+1]] = 5

			counter += 2
		}else if opCode == 4 {
			//output
			if param1PositionMode {
				fmt.Println(input[input[counter+1]])
			}else{
				fmt.Println(input[counter+1])
			}

			counter += 2
		}else if opCode == 5 || opCode == 6 {
			param1 := input[counter+1]
			if param1PositionMode {
				param1 = input[input[counter+1]]
			}

			param2 := input[counter+2]
			if param2PositionMode {
				param2 = input[input[counter+2]]
			}

			if (opCode == 5 && param1 != 0)||(opCode == 6 && param1 == 0) {
				counter = param2
			} else {
				counter += 3
			}
		} else if opCode == 7 || opCode == 8 {
			param1 := input[counter+1]
			if param1PositionMode {
				param1 = input[input[counter+1]]
			}

			param2 := input[counter+2]
			if param2PositionMode {
				param2 = input[input[counter+2]]
			}

			if (opCode == 7 && param1 < param2) || (opCode == 8 && param1 == param2) {
				input[input[counter+3]] = 1
			} else {
				input[input[counter+3]] = 0
			}

			counter += 4
		}


	}

	fmt.Println(input)

	return input[0]
}