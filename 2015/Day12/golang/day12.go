package main

import (
	"regexp"
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	input := ReadFile("prod.txt")
	nums := AllNumbers(input)
	sum := SumNumber(nums)
	fmt.Println(sum)
	nums = AllNumbersIgnoreRed(input)
	sum = SumNumber(nums)
	fmt.Println(sum)
	//fmt.Println(input)
	//50636 too low
	//60709 too low
	//89979 too high
	//63382 wrong
	//65301
	//91907 
}

func ReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return string(b)
}

func AllNumbers(input string) []int {
	numberReg := `\-{0,1}[0-9]+`
	r := regexp.MustCompile(numberReg)
	strResults := r.FindAllString(input,-1)
	intResults := make([]int,0)
	for _, input := range strResults {
		num, _ := strconv.ParseInt(input,10,32)
		intResults = append(intResults,int(num))
	}
	return intResults
}

func AllNumbersIgnoreRed(input string) []int {
	substr, ok := RedSubstr(input)
	for ok {
		if (substr == `:"red"`) {
			fmt.Println("WHOA ", input)
		}
		input = strings.Replace(input,substr,`"XXX"`,1)
		substr, ok = RedSubstr(input)
	}
	fmt.Println("FINAL BEGIN")
	fmt.Println(input)
	fmt.Println("FINAL END")
	return AllNumbers(input)
}

func RedSubstr(input string) (string, bool) {
	left := strings.Index(input,`:"red"`)
	right := left + 6
	if left == -1 {
		return input, false
	}
	counter := 1
	for i := left; i >= 0; i-- {
		if input[i] == '}' {
			counter++
		} else if input[i] == '{' {
			counter--
		} 
		if counter == 0 {
			left = i
			break
		}
	}
	if (counter != 0) {
		fmt.Println("LEFT ERROR ",input)
	}
	
	counter = 1
	for i := right; i < len(input); i++ {
		if input[i] == '}' || input[i] == ']' {
			counter--
		} else if input[i] == '{' || input[i] == '[' {
			counter++
		}
		if counter == 0 {
			right = i+1
			break
		}
	}
	if (counter != 0) {
		fmt.Println("RIGHT ERROR ",input)
	}
	// fmt.Println(string(input[:left]))
	fmt.Println(string(input[left:right]))
	// fmt.Println(string(input[right:len(input)]))
	return input[left:right], true
}

func SumNumber(inputs []int) int {
	count := 0
	for _, input := range inputs {
		count += input
	}
	return count
}