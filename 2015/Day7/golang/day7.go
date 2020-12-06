package main

import (
	"regexp"
	"strconv"
	"fmt"
	"io/ioutil"
	"strings"
)

var cachedValues map[string]uint16

func main() {
	cachedValues = make(map[string]uint16)
	instructions := Instructions(ReadFile("prod.txt"))
	a := GetValue("a",instructions)
	fmt.Println(a)
	cachedValues = make(map[string]uint16)
	cachedValues["b"] = a
	aPrime := GetValue("a",instructions)
	fmt.Println(aPrime)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func Instructions(inputs []string) map[string]string {
	instructions := make(map[string]string)
	for _, input := range inputs {
		tokens := strings.Split(input," -> ")
		instructions[tokens[1]] = tokens[0]
	}
	return instructions
}

func GetValue(input string, instructions map[string]string) uint16 {
	if value, ok := cachedValues[input]; ok {
		return value
	}
	andString := "^(.+) AND (.+)$"
	rAnd := regexp.MustCompile(andString)
	orString := "^(.+) OR (.+)$"
	rOr := regexp.MustCompile(orString)
	lshiftString := "^(.+) LSHIFT ([0-9]+)$"
	rLshift := regexp.MustCompile(lshiftString)
	rshiftString := "^(.+) RSHIFT ([0-9]+)$"
	rRshift := regexp.MustCompile(rshiftString)
	notString := "^NOT (.+)$"
	rNot := regexp.MustCompile(notString)
	valueString := "^([0-9]+)$"
	rValue := regexp.MustCompile(valueString)
	variableString := "^([a-zA-Z]+)$"
	rVariable := regexp.MustCompile(variableString)

	instruction := instructions[input]
	//fmt.Println(instruction," -> ",input)
	switch {
	case rAnd.MatchString(instruction):
		tokens := rAnd.FindStringSubmatch(instruction)
		var a uint16 = 0
		var b uint16 = 0
		if rValue.MatchString(tokens[1]) {
			aPrime,_ := strconv.ParseInt(tokens[1],10,32)
			a = uint16(aPrime)
		} else {
			a = GetValue(tokens[1],instructions)
		}
		if rValue.MatchString(tokens[2]) {
			bPrime,_ := strconv.ParseInt(tokens[2],10,32)
			b = uint16(bPrime)
		} else {
			b = GetValue(tokens[2],instructions)
		}
		cachedValues[input] = a & b
		return a & b
	case rOr.MatchString(instruction):
		tokens := rOr.FindStringSubmatch(instruction)
		var a uint16 = 0
		var b uint16 = 0
		if rValue.MatchString(tokens[1]) {
			aPrime,_ := strconv.ParseInt(tokens[1],10,32)
			a = uint16(aPrime)
		} else {
			a = GetValue(tokens[1],instructions)
		}
		if rValue.MatchString(tokens[2]) {
			bPrime,_ := strconv.ParseInt(tokens[2],10,32)
			b = uint16(bPrime)
		} else {
			b = GetValue(tokens[2],instructions)
		}
		cachedValues[input] = a | b
		return a | b
	case rLshift.MatchString(instruction):
		tokens := rLshift.FindStringSubmatch(instruction)
		a := GetValue(tokens[1],instructions)
		bPrime,_ := strconv.ParseInt(tokens[2],10,32)
		b := uint16(bPrime)
		cachedValues[input] = a << b
		return a << b
	case rRshift.MatchString(instruction):
		tokens := rRshift.FindStringSubmatch(instruction)
		a := GetValue(tokens[1],instructions)
		bPrime,_ := strconv.ParseInt(tokens[2],10,32)
		b := uint16(bPrime)
		cachedValues[input] = a >> b
		return a >> b
	case rNot.MatchString(instruction):
		tokens := rNot.FindStringSubmatch(instruction)
		a := GetValue(tokens[1],instructions)
		cachedValues[input] = ^a
		return ^a
	case rValue.MatchString(instruction):
		prime,_ := strconv.ParseInt(instruction,10,32)
		cachedValues[input] = uint16(prime)
		return uint16(prime)
	case rVariable.MatchString(instruction):
		cachedValues[input] = GetValue(instruction,instructions)
		return cachedValues[input]
	default:
		fmt.Println("Error parsing: ",instruction)
		cachedValues[input] = 0
		return 0
	}

}

