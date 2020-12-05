package main

import (
	"testing"
	//"strconv"
	"fmt"
	"reflect"
)

func TestGetHouses(t *testing.T) {
	inputs := []string{">","^>v<","^v^v^v^v^v"}
	want := []int{2,4,2}
	got := make([]int,0)
	for _, input := range inputs {
		got = append(got,len(GetHouses(input)))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("GetHouses(%q) FAIL",inputs)
		fmt.Println(got)
		fmt.Println(want)
   }
}

func TestGetHousesWithRobot(t *testing.T) {
	inputs := []string{"^v","^>v<","^v^v^v^v^v"}
	want := []int{3,3,11}
	got := make([]int,0)
	for _, input := range inputs {
		got = append(got,len(GetHousesWithRobot(input)))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("GetHousesWithRobot(%q) FAIL",inputs)
		fmt.Println(got)
		fmt.Println(want)
   }
}