package main

import (
	"testing"
	"fmt"
	//"reflect"
	"strconv"
	//"sort"
	"strings"
)

func TestGetMinWait(t *testing.T) {
	lines := ReadFile("test.txt")
	target, _ := strconv.ParseInt(lines[0],10,32)
	trains := GetTimes(strings.Split(lines[1],","))
	index,wait := GetMinWait(target,trains)
	fmt.Println(target,trains,wait)
	var want int64 = 295
	if got := wait * trains[index]; want != got {
		t.Errorf("GetMinWait(text.txt) %d, want %d", got ,want)
	}
}

func TestGetMinWait2(t *testing.T) {
	lines := ReadFile("test.txt")
	target, _ := strconv.ParseInt(lines[0],10,32)
	buses := GetTimesWithOffsets(strings.Split(lines[1],","))
	fmt.Println(target,buses)
	var want int64 = 1068781
	if got := GetFirstIntersection(buses); want != got {
		t.Errorf("GetMinWait(text.txt) %d, want %d", got ,want)
	}
}

func TestGetMinWaitP(t *testing.T) {
	lines := ReadFile("prod.txt")
	target, _ := strconv.ParseInt(lines[0],10,32)
	buses := GetTimesWithOffsets(strings.Split(lines[1],","))
	fmt.Println(target,buses)
	var want int64 = 1068781
	if got := GetFirstIntersection(buses); want != got {
		t.Errorf("GetMinWait(text.txt) %d, want %d", got ,want)
	}
}