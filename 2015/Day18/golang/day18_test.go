package main

import (
	"testing"
	//"fmt"
	"reflect"
)

func TestFrames(t *testing.T) {
	mapl := LoadMap(ReadFile("test.txt"))
	
	want := []int{15,11,8,4,4}
	got := make([]int,0)
	for i := 0; i < 5; i++ {
		got = append(got, CountOn(mapl))
		//fmt.Println(mapl)
		mapl = NextFrame(mapl)
	}
	if !reflect.DeepEqual(want,got) {
		t.Error("NextFrame(test.txt)")
		t.Error(got)
		t.Error(want)
	}
}

func TestFrames2(t *testing.T) {
	mapl := LoadMap(ReadFile("test.txt"))
	
	want := []int{18,18,18,14,17}
	got := make([]int,0)
	for i := 0; i < 5; i++ {
		
		//fmt.Println(mapl)
		mapl = NextFrame2(mapl)
		got = append(got, CountOn(mapl))
	}
	if !reflect.DeepEqual(want,got) {
		t.Error("NextFrame(test.txt)")
		t.Error(got)
		t.Error(want)
	}
}

// func TestNeighbors(t *testing.T) {
// 	mapl := LoadMap(ReadFile("test.txt"))
// 	for i, row := range mapl {
// 		for j, _ := range row {
// 			fmt.Println(GetNeighbors(Coordinate { x: i, y: j},mapl))
// 		}
// 	}
// }