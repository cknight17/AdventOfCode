package main

import (
	"testing"
	"strconv"
	"reflect"
)

func TestUnquote(t *testing.T) {
	inputs := []string{`""`,`"abc"`,`"aaa\"aaa"`,`"\x27"`}
	wantsRaw := []int{2,5,10,6}
	wantsQuoted := []int{0,3,7,1}
	gotRaw := make([]int,0)
	gotQuoted := make([]int,0)
	for _, input := range inputs {
		gotRaw = append(gotRaw,len(input))
		quotedInput,_ := strconv.Unquote(input)
		gotQuoted = append(gotQuoted,len(quotedInput))
	}
	if !reflect.DeepEqual(wantsRaw,gotRaw) || !reflect.DeepEqual(wantsQuoted,gotQuoted) {
		t.Error("TestUnquote()")
		t.Error(gotRaw,wantsRaw)
		t.Error(gotQuoted,wantsQuoted)
	}
}

func TestUnquote2(t *testing.T) {
	inputs := ReadFile("test.txt")
	want:= 12
	got := 0
	for _, input := range inputs {
		got += len(input)
		quotedInput,_ := strconv.Unquote(input)
		got = got - len(quotedInput)
	}
	if want != got {
		t.Error("TestUnquote() ", got," want ", want)
	}
}

func TestQuote2(t *testing.T) {
	inputs := ReadFile("test.txt")
	want:= 19
	got := 0
	for _, input := range inputs {
		got = got - len(input)
		quotedInput := strconv.Quote(input)
		got = got + len(quotedInput)
	}
	if want != got {
		t.Error("TestUnquote() ", got," want ", want)
	}
}