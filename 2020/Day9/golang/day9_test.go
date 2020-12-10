package main

import (
	"testing"
	//"fmt"
	"reflect"
)

func TestSets(t *testing.T) {
	preamble := ReadFile("test.txt")
	sets := Sets(preamble,2)
	values := ValueMap(sets)
	if sets == nil || values == nil {
		t.Error("NO SET")
	}
	inputs := []int64{26,49,100,50}
	want := []bool{true,true,false,false}
	got := make([]bool,0)
	for _, input := range inputs {
		found, ok := values[input]
		got = append(got, found && ok)
	}
	if !reflect.DeepEqual(want,got) {
		t.Error("ValueMap(1-25)")
		t.Error(got)
		t.Error(want)
	}
}

func TestLookback(t *testing.T) {
	set := ReadFile("test2.txt")
	var want int64 = 127
	got, ok := FindFirstFail(set, 5, 2)
	if !ok || want != got {
		t.Errorf("FindFirstFail(test2.txt,5,2) %d, want %d",got,want)
	}
	validRange, _ := ContiguousRange(set,got)
	//fmt.Println(validRange)
	want = 62
	got2 := validRange[0] + validRange[len(validRange)-1]
	if want != got2 {
		t.Errorf("ContiguousRange(test2.txt,%d) %d, want %d",got,got2,want)
	}
}