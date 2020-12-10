package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	inputs := LoadInput(ReadFile("prod.txt"))
	got := FindDiffs(inputs)
	// 1820
	fmt.Println(got,got[1] * got[3])
	fmt.Println(NumCombos(FindGaps(inputs)))
	

}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func LoadInput(input []string) []int {
	output := []int{0}
	for _, item := range input {
		value, _ := strconv.ParseInt(item,10,32)
		output = append(output,int(value))
	}
	sort.Ints(output)
	output = append(output,output[len(output)-1]+3)
	return output
}

func FindDiffs(inputs []int) map[int]int {
	valCounter := map[int]int{1:0,2:0,3:0}
	for i, input := range inputs {
		if i != 0 {
			lastVal := inputs[i-1]
			valCounter[input-lastVal]++
		}
	}
	return valCounter
}

func FindGaps(inputs []int) [][]int {
	gaps := make([][]int,0)
	contig := make([]int,0)

	for index, value := range inputs {
		if index != len(inputs) - 1 {
			nextVal := inputs[index+1]
			diff := nextVal - value
			if diff < 3 {
				contig 	= append(contig,index)
			} else if diff == 3 && len(contig) > 0 {
				contig = append(contig,index)
				gaps = append(gaps,contig)
				contig = make([]int,0)
			}
		}
	}
	return gaps
}

func NumCombos(gaps [][]int) int64 {
	var accumulator int64 = 1
	lookup := map[int]int{5:7,4:4,3:2,2:1,1:1}
	for _, gap := range gaps {
		value, ok := lookup[len(gap)]
		if ok {
			accumulator *= int64(value)
		} else {
			fmt.Printf("ERROR %d NOT FOUND \n",len(gap))
			fmt.Println(gap)
		}
	}
	return accumulator
}