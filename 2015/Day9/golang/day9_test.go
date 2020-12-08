package main

import (
	"testing"
)

func TestFiles(t *testing.T) {
	locations, _ := LoadPaths(ReadFile("test.txt"))
	want := 3
	if got := len(locations); want != got {
		t.Errorf("LoadPaths(test.txt) %d, want %d",got,want)
	}
}

func TestPermutations(t *testing.T) {
	locations, _ := LoadPaths(ReadFile("test.txt"))
	want := 6
	if got := len(AllPaths(locations)); want != got {
		t.Errorf("permutations(test.txt) %d, want %d", got, want)
	}
}

func TestMinPath(t *testing.T) {
	locations, routes := LoadPaths(ReadFile("test.txt"))
	want := 605
	if path, got := MinPath(locations,routes); want != got {
		t.Errorf("MinPath(test.txt) %d, want %d",got,want)
		t.Error(path)
	}
}

func TestMaxPath(t *testing.T) {
	locations, routes := LoadPaths(ReadFile("test.txt"))
	want := 982
	if path, got := MaxPath(locations,routes); want != got {
		t.Errorf("MinPath(test.txt) %d, want %d",got,want)
		t.Error(path)
	}
}