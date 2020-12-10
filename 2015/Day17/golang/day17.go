package main

import (
	"github.com/mxschmitt/golang-combinations"
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	sets := Sets(ReadFile("prod.txt"))
	sets = SetsTarget(150,sets)
	fmt.Println(sets)
	fmt.Println("TOTAL ",len(sets))
	sets = FindMins(sets)
	fmt.Println(sets)
	fmt.Println("TOTAL MIN ",len(sets))
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func Sets(set []string) [][]string {
	return combinations.All(set)
}

func SetsTarget(target int,sets [][]string) [][]string {
	newset := make([][]string,0)
	for _, set := range sets {
		if SumItems(set) == target {
			newset = append(newset,set)
		}
	}
	return newset
}

func FindMins(sets [][]string) [][]string {
	mins := make([][]string,0)
	minVal := len(sets[0])
	for _, set := range sets {
		if len(set) < minVal {
			minVal = len(set)
		}
	}
	for _, set := range sets {
		if len(set) == minVal {
			mins = append(mins,set)
		}
	}
	return mins
}

func SumItems(items []string) int {
	total := 0
	for _, item := range items {
		value, _ := strconv.ParseInt(item,10,32)
		total += int(value)
	}
	return total
}