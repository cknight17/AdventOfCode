package main

import (
	"testing"
	"reflect"
)

func TestThreeVowels(t *testing.T) {
	inputs := []string{"ugknbfddgicrmopn","aaa","jchzalrnumimnmhp","haegwjzuvuyypxyu","dvszwmarrgswjxmb"}
	want := []bool{true,true,true,true,false}
	got := make([]bool,0)
	for _, input := range inputs {
		got = append(got,ThreeVowels(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("TestThreeVowels\n")
		t.Error(inputs)
		t.Error(got)
		t.Error(want)
	}
}

func TestDupeLetter(t *testing.T) {
	inputs := []string{"ugknbfddgicrmopn","aaa","jchzalrnumimnmhp","haegwjzuvuyypxyu","dvszwmarrgswjxmb"}
	want := []bool{true,true,false,true,true}
	got := make([]bool,0)
	for _, input := range inputs {
		got = append(got,DupeLetter(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("TestDupeLetter\n")
		t.Error(inputs)
		t.Error(got)
		t.Error(want)
	}
}

func TestNotAllowed(t *testing.T) {
	inputs := []string{"ugknbfddgicrmopn","aaa","jchzalrnumimnmhp","haegwjzuvuyypxyu","dvszwmarrgswjxmb"}
	want := []bool{true,true,true,false,true}
	got := make([]bool,0)
	for _, input := range inputs {
		got = append(got,NotAllowed(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("NotAllowed\n")
		t.Error(inputs)
		t.Error(got)
		t.Error(want)
	}
}

func TestNiceString(t *testing.T) {
	inputs := []string{"ugknbfddgicrmopn","aaa","jchzalrnumimnmhp","haegwjzuvuyypxyu","dvszwmarrgswjxmb"}
	want := []bool{true,true,false,false,false}
	got := make([]bool,0)
	for _, input := range inputs {
		got = append(got,NiceString(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("NiceString\n")
		t.Error(inputs)
		t.Error(got)
		t.Error(want)
	}
}

func TestDupePair(t *testing.T) {
	inputs := []string{"qjhvhtzxzqqjkmpb","xxyxx","uurcxstgmygtbstg","ieodomkazucvgmuy","aaa"}
	want := []bool{true,true,true,false,false}
	got := make([]bool,0)
	for _, input := range inputs {
		got = append(got,DupePair(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("DupePair\n")
		t.Error(inputs)
		t.Error(got)
		t.Error(want)
	}
}

func TestTriple(t *testing.T) {
	inputs := []string{"qjhvhtzxzqqjkmpb","xxyxx","uurcxstgmygtbstg","ieodomkazucvgmuy"}
	want := []bool{true,true,false,true}
	got := make([]bool,0)
	for _, input := range inputs {
		got = append(got,Triple(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("Triple\n")
		t.Error(inputs)
		t.Error(got)
		t.Error(want)
	}
}

func TestNiceString2(t *testing.T) {
	inputs := []string{"qjhvhtzxzqqjkmpb","xxyxx","uurcxstgmygtbstg","ieodomkazucvgmuy"}
	want := []bool{true,true,false,false}
	got := make([]bool,0)
	for _, input := range inputs {
		got = append(got,NiceString2(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("NiceString2\n")
		t.Error(inputs)
		t.Error(got)
		t.Error(want)
	}
}