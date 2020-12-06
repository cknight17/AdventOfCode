package main

import (
	"regexp"
	"strings"
	"fmt"
	"io/ioutil"
	"math"
)

func main() {
	inputs := ReadFile("prod.txt")
	counter := 0
	counter2 := 0
	for _, input := range inputs {
		if NiceString(input) {
			counter++
		}
		if NiceString2(input) {
			counter2++
		}
	}
	fmt.Println(counter)
	fmt.Println(counter2)
	// 56 wrong
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func ThreeVowels(input string) bool {
	vowelFormat := "[aeiouAEIOU]"
	r, _ := regexp.Compile(vowelFormat)
	return len(r.FindAllString(input,-1)) >= 3
}

func DupeLetter(input string) bool {
	dupeFormat := "aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz|AA|BB|CC|DD|EE|FF|GG|HH|II|JJ|KK|LL|MM|NN|OO|PP|QQ|RR|SS|TT|UU|VV|WW|XX|YY|ZZ"
	r, _ := regexp.Compile(dupeFormat)
	return len(r.FindAllString(input,-1)) >= 1
}

func NotAllowed(input string) bool {
	notAllowedFormat := "ab|cd|pq|xy"
	r, _ := regexp.Compile(notAllowedFormat)
	return len(r.FindAllString(input,-1)) == 0
}

func NiceString(input string) bool {
	return ThreeVowels(input) && DupeLetter(input) && NotAllowed(input)
}

func DupePair(input string) bool {
	for at := 0; at < len(input) - 1; at++ {
		currentPair := []byte{input[at],input[at+1]}
		for look := 0; look < len(input) - 1; look++ {
			if math.Abs(float64(look - at)) >= 2 {
				lookPair := []byte{input[look],input[look+1]}
				//fmt.Printf("Compare pairs %q %q at:%d look:%d\n",currentPair,lookPair,at,look)
				if currentPair[0] == lookPair[0] && currentPair[1] == lookPair[1] {
					return true
				}
			}
		}
	}
	return false
}

func Triple(input string) bool {
	for at := 0; at < len(input) - 2; at++ {
		currentString := []byte{input[at],input[at+1],input[at+2]}
		if currentString[0] == currentString[2] && currentString[0] != currentString[1] {
			return true
		}
	}
	return false
}

func NiceString2(input string) bool {
	return DupePair(input) && Triple(input)
}