package main

import (
	"strconv"
	"fmt"
	"math"
)

func main() {
	input := "3113322113"
	fmt.Println(input)
	for i := 0; i < 40; i++ {
		input = LookAndSay(input)
		//fmt.Println(input)
		fmt.Println(i)
	}
	fmt.Println(len(input))
	input = "3113322113"
	for i := 0; i < 50; i++ {
		input = LookAndSay(input)
		//fmt.Println(input)
		fmt.Println(i)
	}
	fmt.Println(len(input))
}

func LookAndSay(input string) string {
	outputString := ""
	lastChar := input[0]
	lastCounter := 1
	for i := 1; i < len(input); i++ {
		currentChar := input[i]
		if currentChar != lastChar {
			outputString += strconv.Itoa(lastCounter) + string(lastChar)
			lastChar = currentChar
			lastCounter = 1
		} else {
			lastCounter++
		}
	}
	outputString += strconv.Itoa(lastCounter) + string(lastChar)
	return outputString
}

func LookAndSayApprox(input string, power int) int64 {
	conway := 1.303577269
	multiplier := math.Pow(conway,float64(power))
	return int64(len(input)) * int64(multiplier)
}