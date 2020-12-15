package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
	//"math"
)

type Computer struct {
	mask string
	previousMasks []string
	registry map[int64]int64
}

func main() {
	c := Computer{
		mask:"", 
		previousMasks:make([]string,0),
		registry:make(map[int64]int64,0),
	}
	for _, line := range ReadFile("prod.txt") {
		c = ParseInstruction(line,c)
		//fmt.Println(c)
	}
	var sum int64 = 0
	for _, value := range c.registry {
		sum += value
	}
	fmt.Println(sum)

	c = Computer{
		mask:"", 
		previousMasks:make([]string,0),
		registry:make(map[int64]int64,0),
	}
	for _, line := range ReadFile("prod.txt") {
		c = ParseInstruction2(line,c)
		//fmt.Println(c)
	}
	sum = 0
	for _, value := range c.registry {
		sum += value
	}
	fmt.Println(sum)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func BInt(input int64) string {
	base := strconv.FormatInt(input,2)
	//fmt.Println(32-len(base))
	pad := len(base)
	for i := 0; i < 36-pad; i++ {
		base = string('0') + base
		//fmt.Println(base,i)
	}
	//fmt.Println(len(base))
	return base
}

func ProcesssInt(input string,mask string) int64 {
	inputInt,_ := strconv.ParseInt(input,10,64)
	buffer := []rune(BInt(inputInt))
	for index,value := range mask {
		if value != 'X' {
			buffer[index] = value
		}
	}
	bufferInt,_ := strconv.ParseInt(string(buffer),2,64)
	return bufferInt
}

func ProcesssInt2(input string,mask string) []int64 {
	inputInt,_ := strconv.ParseInt(input,10,64)
	buffer := []rune(BInt(inputInt))
	//obuffer := string(buffer)
	
	for index,value := range mask {
		if value == 'X' {
			buffer[index] = 'X'
		} else if value == '1' {
			buffer[index] = '1'
		}
	}
	//fmt.Println(obuffer,mask,string(buffer))
	x := strings.Count(string(buffer),"X")
	binStrs := getBinaryOptions(make([]string,0),x)
	//fmt.Println("BIN",binStrs,string(buffer),x)
	bufferStrs := make([]string,0)
	for _, binStr := range binStrs {
		tmp := string(buffer)
		for i := 0; i < len(binStr); i++ {
			tmp = strings.Replace(tmp,"X",string(binStr[i]),1)
		}
		bufferStrs = append(bufferStrs,tmp)
	}
	
	//fmt.Println(bufferStrs)
	bufferInts := make([]int64,0)
	for _, bufferStr := range bufferStrs {
		bufferInt,_ := strconv.ParseInt(string(bufferStr),2,64)
		bufferInts = append(bufferInts,bufferInt)
	}
	//fmt.Println(bufferInts)
	return bufferInts
}

func getBinaryOptions(input []string,num int) []string {
	if num  == 0 {
		return input
	} else if len(input) == 0 {
		return getBinaryOptions([]string{"0","1"},num-1)
	}
	newBatch := make([]string,0)
	for _, item := range input {
		newBatch = append(newBatch,"0" + item)
	}
	for _, item := range input {
		newBatch = append(newBatch,"1" + item)
	}
	return getBinaryOptions(newBatch,num-1)
}

func ParseInstruction(input string, c Computer) Computer {
	if strings.HasPrefix(input,"mem") {
		//fmt.Println("mem")
		regStr := `^mem\[([0-9]+)\] \= ([0-9]+)$`
		//fmt.Println(input)
		r := regexp.MustCompile(regStr)
		match := r.FindStringSubmatch(input)
		//fmt.Println(match)
		index,_ := strconv.ParseInt(match[1],10,64)
		c.registry[index] = ProcesssInt(match[2],c.mask)
	} else {
		//fmt.Println("mask")
		if c.mask != "" {
			c.previousMasks = append(c.previousMasks,c.mask)
		}
		c.mask = input[7:]
	}
	return c
}

func ParseInstruction2(input string, c Computer) Computer {
	if strings.HasPrefix(input,"mem") {
		//fmt.Println("mem")
		regStr := `^mem\[([0-9]+)\] \= ([0-9]+)$`
		//fmt.Println(input)
		r := regexp.MustCompile(regStr)
		match := r.FindStringSubmatch(input)
		//fmt.Println(match)
		index,_ := strconv.ParseInt(match[2],10,64)
		ints := ProcesssInt2(match[1],c.mask)
		for i := 0; i < len(ints); i++ {
			c.registry[ints[i]] = index
		}
	} else {
		//fmt.Println("mask")
		if c.mask != "" {
			c.previousMasks = append(c.previousMasks,c.mask)
		}
		c.mask = input[7:]
	}
	return c
}

