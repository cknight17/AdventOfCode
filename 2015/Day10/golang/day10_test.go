package main

import (
	"testing"
	"reflect"
)

func TestLookAndSay(t *testing.T) {
	want := []string{"11","21","1211","111221","312211"}
	inputs := []string{"1","11","21","1211","111221"}
	got := make([]string,0)
	for _, input := range inputs {
		got = append(got,LookAndSay(input))
	}
	if !reflect.DeepEqual(want,got) {
		t.Error("LookAndSay(",inputs,") ",got," want ",want)
	}
}

func TestApprox(t *testing.T) {
	want := 329356
	input := "3113322113"
	if got := LookAndSayApprox(input,39); int64(want) != got {
		t.Errorf("APprox(%q) %d %d",input,got,want)
	}
}