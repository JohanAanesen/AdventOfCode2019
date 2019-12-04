package main

import (
	"fmt"
	"os"
)

func main() {

	min := 245318
	max := 765757

	counter := 0
	counter2 := 0

	for i := min; i < max; i++ {

		password := true

		double := false

		nr1 := (i / 100000) % 10
		nr2 := (i / 10000) % 10
		nr3 := (i / 1000) % 10
		nr4 := (i / 100) % 10
		nr5 := (i / 10) % 10
		nr6 := i % 10

		if nr1 > nr2 {
			password = false
		} else if nr1 == nr2 {
			double = true
		}

		if nr2 > nr3 {
			password = false
		} else if nr2 == nr3 {
			double = true
		}

		if nr3 > nr4 {
			password = false
		} else if nr3 == nr4 {
			double = true
		}

		if nr4 > nr5 {
			password = false
		} else if nr4 == nr5 {
			double = true
		}

		if nr5 > nr6 {
			password = false
		} else if nr5 == nr6 {
			double = true
		}

		if password && double {
			//fmt.Println(i)
			counter++
		}


		//part 2
		count := map[int]int{}
		count[nr1]++
		count[nr2]++
		count[nr3]++
		count[nr4]++
		count[nr5]++
		count[nr6]++


		if password {
			for _, v := range count {
				if v == 2 {
					counter2++
					break
				}
			}
		}

	}

	//part 1
	fmt.Println(counter)
	//part 2
	fmt.Println(counter2)

	os.Exit(0) //quit

}
