package main

import (
	"testing"
	"reflect"
	//"fmt"
)

func TestAllNumbers(t *testing.T) {
	inputs := []string{`[1,2,3]`,`{"a":2,"b":4}`,`[[[3]]]`,`{"a":{"b":4},"c":-1}`,`{"a":[-1,1]}`,`[-1,{"a":1}]`,`[]`,`{}`}
	wants := []int{6,6,3,3,0,0,0,0}
	gots := make([]int,0)
	for _, input := range inputs {
		nums := AllNumbers(input)
		gots = append(gots,SumNumber(nums))
	}
	if !reflect.DeepEqual(wants,gots) {
		t.Errorf("AllNumbers(SumNumber(%q))",inputs)
		t.Error(gots," want ", wants)
	}
}

func TestAllNumbersIR(t *testing.T) {
	inputs := []string{`[1,2,3]`,`[1,{"c":"red","b":2},3]`,`{"d":"red","e":[1,2,3,4],"f":5}`,`[1,"red",5]`,`{"a":1,"b":{"c":1,"d":"red","e":{"f":1}},"g":[1,2,3,"red"],"h":{"i":"red"}}`}
	wants := []int{6,4,0,6,7}
	gots := make([]int,0)
	for _, input := range inputs {
		nums := AllNumbersIgnoreRed(input)
		gots = append(gots,SumNumber(nums))
	}
	if !reflect.DeepEqual(wants,gots) {
		t.Errorf("AllNumbers(SumNumber(%q))",inputs)
		t.Error(gots," want ", wants)
	}
}

func TestRedSubstr(t *testing.T) {
	input := `{"a":1,"b":{"c":1,"d":"red","e":{"f":1}},"g":[1,2,3,"red"],"h":{"i":"red"}}`
	RedSubstr(input)
}
