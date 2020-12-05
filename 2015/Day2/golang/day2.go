package main

import (
	"strconv"
	"io/ioutil"
	"fmt"
	"strings"
	"sort"
)

func main() { 
	inputs := ReadFile("prod.txt")
	totalArea := 0
	totalLength := 0
	for _, input := range inputs {
		totalArea += PackageArea(ParseDimensions(input))
		totalLength += RibbonLength(ParseDimensions(input))
	} 
	fmt.Println(totalArea)
	fmt.Println(totalLength)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func ParseDimensions(input string) []int {
	values := make([]int,0)
	strs := strings.Split(input,"x")
	for _, strnum := range strs {
		value, _ := strconv.ParseInt(strnum,10,32)
		values = append(values,int(value))
	}
	return values
}

func PackageArea(input []int) int {
	sort.Ints(input)
	if (len(input) != 3) {
		fmt.Printf("ERROR: expecting 3, got %d\n",len(input))
		return -1
	}
	return (input[0] * input[1] * 3) + (input[1] * input[2] * 2) + (input[0] * input[2] * 2)
}

func RibbonLength(input []int) int {
	sort.Ints(input)
	if (len(input) != 3) {
		fmt.Printf("ERROR: expecting 3, got %d\n",len(input))
		return -1
	}
	//fmt.Printf("CALC: (%d * 2) + (%d * 2) + (%d * %d * %d) = %d\n",input[0],input[1],input[0],input[1],input[2],(input[0] * 2) + (input[1] * 2) + (input[0] * input[1] * input[2]))
	return (input[0] * 2) + (input[1] * 2) + (input[0] * input[1] * input[2])
}