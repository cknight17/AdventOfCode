package main

import (
	"testing"
	//"fmt"
	"reflect"
)

func TestLoad(t *testing.T) {
	input := ReadFile("test.txt")
	
	input2 := NextFrame(input)
	for !reflect.DeepEqual(input,input2) {
		input = input2
		input2 = NextFrame(input)
	}
	counter := 0
	for _, line := range input2 {
		for _, r := range line {
			if r == '#' {
				counter++
			}
		}
	}
	want := 37
	if got := counter; want != got {
		t.Errorf("NextFrame(text.txt) %d, want %d", got ,want)
	}
}

func TestLoad2(t *testing.T) {
	input := ReadFile("test.txt")
	
	input2 := NextFrame2(input)
	for !reflect.DeepEqual(input,input2) {
		input = input2
		input2 = NextFrame2(input)
	}
	counter := 0
	for _, line := range input2 {
		for _, r := range line {
			if r == '#' {
				counter++
			}
		}
	}
	want := 26
	if got := counter; want != got {
		t.Errorf("NextFrame2(text.txt) %d, want %d", got ,want)
	}
}