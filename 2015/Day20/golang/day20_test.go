package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestHouses(t *testing.T) {
	inputs := []int{1,2,3,4,5,6,7,8,9}
	want := []int{10,30,40,70,60,120,80,150,130}
	got := make([]int,0)
	for _, input := range inputs {
		got = append(got,House(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Error("TestHouses(",inputs,") ", got,", want",want)
	}
}

func TestHouse(t *testing.T) {
	want := 36000000
	fmt.Println(2248322)
	f1 := House(2248321)
	f2 := House(2248322)
	f3 := House(2248323)
	fmt.Println(f1, " => ",want-f1)
	fmt.Println(f2, " => ",want-f2)
	fmt.Println(f3, " => ",want-f3)
	// 2248323 too high
}