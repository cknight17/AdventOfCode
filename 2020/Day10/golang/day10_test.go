package main

import (
	"testing"
	"fmt"
	"reflect"
)

func TestPath1(t *testing.T) {
	inputs := LoadInput(ReadFile("test.txt"))
	want := map[int]int{1:7,2:0,3:5}
	if got := FindDiffs(inputs); !reflect.DeepEqual(want,got) {
		t.Error("FIndDiffs(test.txt) ",got," want ",want)
	}
	var want2 int64 = 8
	gaps := FindGaps(inputs)
	num := NumCombos(gaps)
	fmt.Println(inputs)
	fmt.Println(gaps)
	fmt.Println(num)
	if got2 := NumCombos(FindGaps(inputs)); want2 != got2 {
		t.Error("NumCombos(test.txt) ",got2," want ",want2)
	}
}

func TestPath2(t *testing.T) {
	inputs := LoadInput(ReadFile("test2.txt"))
	want := map[int]int{1:22,2:0,3:10}
	if got := FindDiffs(inputs); !reflect.DeepEqual(want,got) {
		t.Error("FIndDiffs(test2.txt) ",got," want ",want)

	}
	var want2 int64 = 8
	gaps := FindGaps(inputs)
	num := NumCombos(gaps)
	fmt.Println(inputs)
	fmt.Println(gaps)
	fmt.Println(num)
	if got2 := NumCombos(FindGaps(inputs)); want2 != got2 {
		t.Error("NumCombos(test.txt) ",got2," want ",want2)
	}
}