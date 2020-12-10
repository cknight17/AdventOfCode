package main

import (
	"github.com/mxschmitt/golang-combinations"
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	set := ReadFile("prod.txt")
	got, ok := FindFirstFail(set, 25, 2);
	fmt.Println(ok,got)
	// 85848519
	validRange, _ := ContiguousRange(set,got)
	got2 := validRange[0] + validRange[len(validRange)-1]
	fmt.Println(got2)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func Sets(set []string, of int) [][]string {
	cmb := combinations.Combinations(set,of)
	return cmb
}

func SumItems(items []string) int64 {
	var total int64 = 0
	for _, item := range items {
		value, _ := strconv.ParseInt(item,10,64)
		total += value
	}
	return total
}

func ValueMap(sets [][]string) map[int64]bool {
	valueMap := make(map[int64]bool,0)
	for _, set := range sets {
		valueMap[SumItems(set)] = true
	}
	return valueMap
}

func ConvertToInt64(set []string) []int64 {
	output := make([]int64,0)
	for _, item := range set {
		val, _ := strconv.ParseInt(item,10,64)
		output = append(output,val)
	}
	return output
}

func LoadPreamble(set []string, size int, at int) []string {
	startp := at - size
	if at < 0 {
		startp = 0
	}
	endp := at
	if at > len(set) {
		endp = len(set)
	}
	return set[startp:endp]
}

func FindFirstFail(set []string, lookBack int, num int) (int64, bool) {
	for i := lookBack; i < len(set); i++ {
		item, _ := strconv.ParseInt(set[i],10,32)
		preamble := LoadPreamble(set,lookBack,i)
		pSets := Sets(preamble,num)
		values := ValueMap(pSets)
		found, ok := values[item]
		if !found || !ok {
			fmt.Println(item,preamble,pSets,values,found,ok)
			return item, true
		}
	}
	return 0, false
}

func ContiguousRange(set []string, target int64) ([]int64, bool) {
	for i := 0; i < len(set); i++ {
		for j := i + 1; j < len(set); j++ {
			currentRange := set[i:j]
			sum := SumItems(currentRange)
			if sum == target {
				validRange := ConvertToInt64(currentRange)
				sort.Slice(validRange, func(i, j int) bool { return validRange[i] < validRange[j] })
				return validRange, true
			} else if sum > target {
				break
			}
		}
	}
	return nil, false
}