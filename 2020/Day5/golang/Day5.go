package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"sort"
)

func main() {
	input := "BFFFBBFRRR"
	input = binaryString(input)
	fmt.Println(strconv.ParseInt(input,2,64))
	input = "FFFBBBFRRR"
	input = binaryString(input)
	fmt.Println(strconv.ParseInt(input,2,64))
	input = "BBFFBBFRLL"
	input = binaryString(input)
	fmt.Println(strconv.ParseInt(input,2,64))
	inputs := ReadFile("prod.txt")
	outputs := make([]int,0)
	for _, input := range inputs {
		value, _ := strconv.ParseInt(binaryString(input),2,64)
		outputs = append(outputs,int(value))
	}
	sort.Ints(outputs)
	fmt.Println(outputs)
	
	fmt.Println(outputs[len(outputs)-1])
	missing := missingSeats(outputs)
	fmt.Println(missing)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func binaryString(input string) string {
	input = strings.ReplaceAll(input,"B","1")
	input = strings.ReplaceAll(input,"F","0")
	input = strings.ReplaceAll(input,"R","1")
	input = strings.ReplaceAll(input,"L","0")
	return input
}

func missingSeats(list []int) []int {
	min := list[0]
	max := list[len(list)-1]
	at := 0
	missing := make([]int,0)
	for i := min; i <= max; i++ {
		if list[at] == i {
			at++
		} else {
			missing = append(missing,i)
		}
	}
	return missing
}

