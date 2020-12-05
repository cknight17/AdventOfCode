package main

import (
	"testing"
)

func TestLeading5(t *testing.T) {
	input := "abcdef"
	var want int64 = 609043
	if got := leading5(input); want != got {
		t.Errorf("leading5(%q) %d, want %d",input,got,want)
   	}
}

func TestLeading5b(t *testing.T) {
	input := "pqrstuv"
	var want int64 = 1048970
	if got := leading5(input); want != got {
		t.Errorf("leading5(%q) %d, want %d",input,got,want)
   	}
}