package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	//"math"
)

func main() {
	lines := ReadFile("prod.txt")
	i64target, _ := strconv.ParseInt(lines[0],10,64)
	target := i64target
	trains := GetTimes(strings.Split(lines[1],","))
	index,wait := GetMinWait(target,trains)
	fmt.Println(target,trains,wait)
	answer := wait * trains[index]; 
	fmt.Println(answer)

	buses := GetTimesWithOffsets(strings.Split(lines[1],","))
	fmt.Println(buses)
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func GetTimes(inputs []string) []int64 {
	outputs := make([]int64,0)
	for _, input := range inputs {
		output, err := strconv.ParseInt(input,10,64)
		if err == nil {
			outputs = append(outputs,output)
		}
	}
	return outputs
}

type Bus struct {
	number int64
	offset int64
}

func GetTimesWithOffsets(inputs []string) []Bus {
	outputs := make([]Bus,0)
	for index, input := range inputs {
		output, err := strconv.ParseInt(input,10,64)
		if err == nil {
			outputs = append(outputs,Bus{number:output,offset:int64(index)})
		}
	}
	return outputs
}

func GetMinWait(target int64,inputs []int64) (int64,int64) {
	var minIndex  int64 = 0
	minWait := GetWait(target,inputs[0])

	for i, input := range inputs {
		wait := GetWait(target,input)
		if wait < minWait {
			minIndex = int64(i)
			minWait = wait
		}
	}
	return minIndex, minWait
}

func GetWait(target int64, input int64) int64 {
	return input - (target % input)
}

func GetFirstIntersection(buses []Bus) int64 {
	//firstBus := buses[0]
	var t int64 = 0
	//var jump int64 = 0
	// for i, bus := range buses {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if i > 1 {
	// 		break
	// 	}
	// 	var j int64 = bus.number
	// 	for {
	// 		if j % firstBus.number == bus.offset {
	// 			fmt.Println("FOUND ",j,bus,firstBus)
	// 			break
	// 		} else {
	// 			fmt.Println("NOT FOUND ",j,bus,firstBus)
	// 		}
	// 		j = j + bus.number
	// 	}
	// 	t = j - bus.offset
	// 	fmt.Println("t: ",t)
	// }
	// t = Increment(0,buses[1].number,buses[1].number,buses[0].number,buses[1].offset)
	t = FirstTwo(buses[0].number,buses[1].number,buses[1].offset)
	LCMSet := []int64{buses[0].number,buses[1].number}
	for _, bus := range buses[2:] {
		t = Increment(t,LCM(LCMSet[0],LCMSet[1],LCMSet[2:]...),bus.number,bus.offset)
		LCMSet = append(LCMSet,bus.number)
	}
	// t = Increment(t,LCM(buses[0].number,buses[1].number),buses[2].number,buses[2].offset)
	// t = Increment(t,LCM(buses[0].number,buses[1].number,buses[2].number),buses[3].number,buses[3].offset)
	// t = Increment(t,LCM(buses[0].number,buses[1].number,buses[2].number,buses[3].number),buses[4].number,buses[4].offset)
	return t
}

func Increment(start int64, jump int64, input int64, offset int64) int64 {
	at := start
	for {
		fmt.Println(at,start,jump,input,offset,at % input)
		if (at+offset) % input == 0 {
			return at
		}
		
		at += jump
	}
	return -1
}

func FirstTwo(x1, x2, offset int64) int64 {
	at := x1
	for {
		fmt.Println("FT",at,x1,x2,x1 % at, (x2+offset)%at)
		if (at+offset) % x2 == 0 {
			return at
		}
		at = at + x1
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
