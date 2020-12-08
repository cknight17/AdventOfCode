package main

import (
	"testing"
	"reflect"
)

func TestNext(t *testing.T) {
	input := "azzzzzzz"
	want := "baaaaaaa"
	if got := NextInput(input); want != got {
		t.Errorf("NextInput(%q) %q, want %q",input,got,want)
	}
}

func TestNextValid(t *testing.T) {
	inputs := []string{"abcdefgh","ghijklmn"}
	wants := []string{"abcdffaa","ghjaabcc"}
	gots := make([]string,0)
	for _, input := range inputs {
		got, ok := NextValid(input)
		if ok {
			gots = append(gots,got)
		}
	}
	if !reflect.DeepEqual(wants,gots) {
		t.Errorf("NextValid(%q)",inputs)
		t.Error(gots," want ",wants)
	}
}

func TestIncreasing(t *testing.T) {
	inputs := []string{"hijklmmn","abbceffg","abbcegjk","abcdffaa","ghjaabcc"}
	wants := []bool{true,false,false,true,true}
	gots := make([]bool,0)
	for _, input := range inputs {
		gots = append(gots,CheckIncreasing(input))
	}
	if !reflect.DeepEqual(wants,gots) {
		t.Errorf("CheckIncreasing(%q)",inputs)
		t.Error(gots," want ",wants)
	}
}

func TestForbidden(t *testing.T) {
	inputs := []string{"hijklmmn","abbceffg","abbcegjk","abcdffaa","ghjaabcc"}
	wants := []bool{false,true,true,true,true}
	gots := make([]bool,0)
	for _, input := range inputs {
		gots = append(gots,CheckForbidden(input))
	}
	if !reflect.DeepEqual(wants,gots) {
		t.Errorf("CheckForbiden(%q)",inputs)
		t.Error(gots," want ",wants)
	}
}

func TestPairs(t *testing.T) {
	inputs := []string{"hijklmmn","abbceffg","abbcegjk","abcdffaa","ghjaabcc"}
	wants := []bool{false,true,false,true,true}
	gots := make([]bool,0)
	for _, input := range inputs {
		gots = append(gots,CheckPairs(input))
	}
	if !reflect.DeepEqual(wants,gots) {
		t.Errorf("CheckPairs(%q)",inputs)
		t.Error(gots," want ",wants)
	}
}

func TestInput(t *testing.T) {
	inputs := []string{"hijklmmn","abbceffg","abbcegjk","abcdffaa","ghjaabcc"}
	wants := []bool{false,false,false,true,true}
	gots := make([]bool,0)
	for _, input := range inputs {
		gots = append(gots,CheckInput(input))
	}
	if !reflect.DeepEqual(wants,gots) {
		t.Errorf("CheckInput(%q)",inputs)
		t.Error(gots," want ",wants)
	}
}