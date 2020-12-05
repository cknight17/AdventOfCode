package main

import (
	"strings"
	"io/ioutil"
	"fmt"
)

func main() { 
	input := ReadFile("prod.txt")
	fmt.Println(Count(input))
	fmt.Println(FindTarget(input,-1))
}

func CountUp(input string) int {
	return strings.Count(input,"(")
}

func CountDown(input string) int {
	return strings.Count(input,")")
}

func FindTarget(input string, target int) int {
	counter := 0
	for index, char := range input {
		if char == '(' {
			counter++
		} else if char == ')' {
			counter--
		}
		if counter == target {
			return index+1
		}
	}
	return 0
}

func ReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return string(b)
}

func Count(input string) int {
	return CountUp(input) - CountDown(input)
}