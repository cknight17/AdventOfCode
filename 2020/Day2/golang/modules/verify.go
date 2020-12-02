package verify

import (
	"aoc.knighttechnology.net/fileParser"
	"strings"
)

func Verify(item fileParser.Day2line) bool {
	count := strings.Count(item.Input,item.Character)
	return item.Min <= count && item.Max >= count
}

func VerifyAll(items []fileParser.Day2line) []bool {
	bools := make([]bool,0)
	for _, item := range items {
		bools = append(bools,Verify(item))
	}
	return bools
}

func VerifyAllCount(items []fileParser.Day2line) int {
	bools := VerifyAll(items)
	counter := 0
	for _, item := range bools {
		if item {
			counter++
		}
	}
	return counter
}

func Verifyb(item fileParser.Day2line) bool {
	isA := item.Input[item.Min-1] == item.Character[0]
	isB := item.Input[item.Max-1] == item.Character[0]
	return isA != isB
}

func VerifyAllb(items []fileParser.Day2line) []bool {
	bools := make([]bool,0)
	for _, item := range items {
		bools = append(bools,Verifyb(item))
	}
	return bools
}


func VerifyAllCountb(items []fileParser.Day2line) int {
	bools := VerifyAllb(items)
	counter := 0
	for _, item := range bools {
		if item {
			counter++
		}
	}
	return counter
}