package main

import (
	"testing"
	"fmt"
	//"reflect"
)

func TestReplace(t *testing.T) {
	blocks := ReadFile("test.txt")
	language := GetReplacement(blocks[0])
	input := blocks[1]
	repls := ReplacementSets(input,language)
	//fmt.Println(language)
	want := 4
	if got := len(repls); want != got {
		t.Errorf("ReplacementSets(test.txt) %d, want %d", got, want)
		t.Error(input)
		t.Error(repls)
	}
}

func TestReplace2(t *testing.T) {
	blocks := ReadFile("test2.txt")
	language := GetReplacement(blocks[0])
	input := blocks[1]
	repls := ReplacementSets(input,language)
	//fmt.Println(language)
	want := 7
	if got := len(repls); want != got {
		t.Errorf("ReplacementSets(test.txt) %d, want %d", got, want)
		t.Error(input)
		t.Error(repls)
	}
}

func TestReplace3(t *testing.T) {
	blocks := ReadFile("test3.txt")
	language := GetReplacement(blocks[0])
	input := blocks[1]
	repls := ReplacementSets(input,language)
	//fmt.Println(language)
	want := 1
	if got := len(repls); want != got {
		t.Errorf("ReplacementSets(test.txt) %d, want %d", got, want)
		t.Error(input)
		t.Error(repls)
	}
}

func TestReplace4(t *testing.T) {
	blocks := ReadFile("test4.txt")
	language := GetReplacement(blocks[0])
	language = Flip(language)
	input := []string{blocks[1]}
	target := "e"
	fmt.Println(MassageInput(input[0]))
	i := 0
	for {
		var found bool
		i++
		//input = ReplacementSetsSets(input,language)
		input, found = Filter(input,target)
		if found {
			break
		}
	}
	//fmt.Println(language)
	want := 3
	if got := i; want != got {
		t.Errorf("ReplacementSetsSets(test.txt) %d, want %d", got, want)
		t.Error(input)
		t.Error(i)
		t.Error(target)
	}
}

// func TestReplace5(t *testing.T) {
// 	blocks := ReadFile("test5.txt")
// 	language := GetReplacement(blocks[0])
// 	//language = Flip(language)
// 	fmt.Println(language)
// 	input := []string{"e"}
// 	target := blocks[1]
// 	fmt.Println(input)
// 	fmt.Println(target)
// 	i := 0
// 	for {
// 		var found bool
// 		i++
// 		input = ReplacementSetsSets(input,language)
// 		input, found = Filter(input,target)
// 		if found {
// 			break
// 		} else if len(input) == 0 {
// 			break
// 		} else {
// 			fmt.Println("NOT FOUND ",input)
// 		}
// 	}
// 	//fmt.Println(language)
// 	want := 6
// 	if got := i; want != got {
// 		t.Errorf("ReplacementSetsSets(test.txt) %d, want %d", got, want)
// 		t.Error(input)
// 		t.Error(i)
// 		t.Error(target)
// 	}
// }

func TestDiff(t *testing.T) {
	input := "ABCDE"
	target := "ABCFFF"
	want1 := "ABC"
	want2 := "DE"
	if got1, got2 := Diff(input,target); want1 != got1 && want2 != got2 {
		t.Errorf("Diff(%q,%q) %q %q, want %q %q",input,target,got1,got2,want1,want2)
	}
}

func TestDiff2(t *testing.T) {
	input := "BBCDE"
	target := "ABCFFF"
	want1 := ""
	want2 := "BBCDE"
	if got1, got2 := Diff(input,target); want1 != got1 && want2 != got2 {
		t.Errorf("Diff(%q,%q) %q %q, want %q %q",input,target,got1,got2,want1,want2)
	}
}