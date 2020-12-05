package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestCount(t *testing.T) {
	inputs := []string {"(())","()()","(((","(()(()(","))(((((","())","))(",")))",")())())"}
	want := []int{0,0,3,3,3,-1,-1,-3,-3}
	got := make([]int,0)
	for _, input := range inputs {
		got = append(got,Count(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("verifyAll(%q) FAIL",inputs)
		fmt.Println(got)
		fmt.Println(want)
   }
}

func TestFindTarget(t *testing.T) {
	inputs := []string {")","()())"}
	want := []int{1,5}
	got := make([]int,0)
	for _, input := range inputs {
		got = append(got,FindTarget(input,-1))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("verifyAll(%q) FAIL",inputs)
		fmt.Println(got)
		fmt.Println(want)
   }
}