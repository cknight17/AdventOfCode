package main

import (
	"aoc.knighttechnology.net/verify"
	"aoc.knighttechnology.net/fileParser"
	"fmt"
)
func main() { 

	file:="day2a.txt"
	output := fileParser.Day2input(file)
	counter := verify.VerifyAllCount(output)
	fmt.Println(counter)

	counterb := verify.VerifyAllCountb(output)
	fmt.Println(counterb)
}