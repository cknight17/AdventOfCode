package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	inputs := ReadFile("prod.txt")
	got := 0
	got2 := 0
	for _, input := range inputs {
		got += len(input)
		got2 = got2 - len(input)
		quotedInput,_ := strconv.Unquote(input)
		got = got - len(quotedInput)
		quotedInput2 := strconv.Quote(input)
		got2 = got2 + len(quotedInput2)
	}
	fmt.Println(got)
	fmt.Println(got2)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}