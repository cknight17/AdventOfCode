package main

import (
	"testing"
	//"fmt"
	"sort"
)

func TestParseHappiness(t *testing.T) {
	input := "Alice would gain 54 happiness units by sitting next to Bob."
	got, got2, got3 := ParseHappiness(input)
	want := "Alice"
	want2 := "Bob"
	want3 := 54
	if want != got || want2 != got2 || want3 != got3 {
		t.Errorf("ParseHappiness(%q) %q %q %d, %q %q %d\n",input,got,got2,got3,want,want2,want3)
	}
}

func TestParseHappiness2(t *testing.T) {
	input := "Carol would lose 83 happiness units by sitting next to Alice."
	got, got2, got3 := ParseHappiness(input)
	want := "Carol"
	want2 := "Alice"
	want3 := -83
	if want != got || want2 != got2 || want3 != got3 {
		t.Errorf("ParseHappiness(%q) %q %q %d, %q %q %d\n",input,got,got2,got3,want,want2,want3)
	}
}

func TestHappinessMap(t *testing.T) {
	inputs := ReadFile("test.txt")
	hmap,_ := HappinessMap(inputs)
	seating := []string{"David","Alice","Bob","Carol"}
	want := 330
	got, ok := CalculateHappiness(seating,hmap)
	if !ok || want != got {
		t.Errorf("HappinessMap(%q) %d, want %d\n",seating,got,want)
	}
}

func TestCalculateHappinessAll(t *testing.T) {
	inputs := ReadFile("test.txt")
	hmap,hlist := HappinessMap(inputs)
	alist := Permutations(hlist)
	allSeating := CalculateHappinessAll(alist,hmap)
	sort.Sort(ByHappiness(allSeating))
	max := allSeating[len(allSeating)-1]
	want := 330
	if got := max.happiness; want != got {
		t.Errorf("CalculateHappinessAll(%q)\n",hlist)
		t.Error(allSeating)
		t.Error(max)
	}
}