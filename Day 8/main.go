package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	x := 25
	y := 6

	pixelsPrImage := x * y

	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	input := ""

	for r.Scan() {
		line := r.Text()


		//#1 @ 342,645: 25x20
		fmt.Sscanf(line, "%s\n", &input)


	}

	fmt.Println(input)

	runeString := []rune(input)

	var stringArr []string

	for i := 0; i < len(input)/pixelsPrImage; i++ {

		tempString := string(runeString[i*pixelsPrImage:(i*pixelsPrImage)+pixelsPrImage])

		stringArr = append(stringArr, tempString)

	}

	highest0 := 999999999
	highestS := ""

	for _, s := range stringArr {

		count0 := strings.Count(s, "0")

		if count0 < highest0 {
			highest0 = count0
			highestS = s
		}
	}
	fmt.Println(strings.Count(highestS, "1") * strings.Count(highestS, "2"))



	var resultImage []rune


	for i := 0; i < 150; i++ {

		for _, s := range stringArr {

			runeS := []rune(s)
			if runeS[i] == '1' || runeS[i] == '0'{
				resultImage = append(resultImage, runeS[i])
				break
			}

		}


	}

	//resultString := string(resultImage)

	for i := 0; i < 6; i++ {
		fmt.Println(string(resultImage[i*25:(i*25)+25]))
	}

}
