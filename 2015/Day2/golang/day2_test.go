package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestReadFile(t *testing.T) {
	want := 1000
	got := len(ReadFile("prod.txt"))
	if want != got {
		t.Errorf("ReadFile() %d, want %d",got,want)
	}
}

func TestParseDimensions(t *testing.T) {
	inputs := []string {"2x3x4","1x1x10"}
	want := [][]int{{2,3,4},{1,1,10}}
	got := make([][]int,0)
	for _, input := range inputs {
		got = append(got,ParseDimensions(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("ParseDimensions(%q) FAIL",inputs)
		fmt.Println(got)
		fmt.Println(want)
   }
}

func TestPackageArea(t *testing.T) {
	inputs := []string { "2x3x4", "1x1x10", "10x1x1" }
	want := []int{ 58, 43, 43 }
	got := make([]int,0)
	for _, input := range inputs {
		got = append(got,PackageArea(ParseDimensions(input)))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("PackageArea(%q) FAIL",inputs)
		fmt.Println(got)
		fmt.Println(want)
   }
}

func TestRibbonLength(t *testing.T) {
	inputs := []string { "2x3x4", "1x1x10", "10x1x1" }
	want := []int{ 34, 14, 14 }
	got := make([]int,0)
	for _, input := range inputs {
		got = append(got,RibbonLength(ParseDimensions(input)))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("RibbonLength(%q) FAIL",inputs)
		fmt.Println(got)
		fmt.Println(want)
   }
}