package main

import (
	"testing"
	//"fmt"
)

func TestSets(t *testing.T) {
	sets := Sets(ReadFile("test.txt"))
	sets = SetsTarget(25,sets)
	want := 4
	if got := len(sets); want != got {
		t.Errorf("SetsTarget(test.txt) %d, want %d\n", got, want)
	}

	sets = FindMins(sets)
	want = 3
	if got := len(sets); want != got {
		t.Errorf("FindMins(test.txt) %d, want %d\n", got, want)
	}
}