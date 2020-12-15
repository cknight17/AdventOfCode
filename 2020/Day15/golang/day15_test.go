package main

import (
	"testing"
	//"reflect"
)

func TestRun(t *testing.T) {
	input := "0,3,6"
	g := LoadMap(input)
	g = RunGame(g,10)
	want := uint64(0)
	if got := g.values[len(g.values)-1]; want != got {
		t.Error("TestFirst(",input,") ", got, " want ", want)
	}
}

func TestRun2(t *testing.T) {
	input := "1,3,2"
	g := LoadMap(input)
	g = RunGame(g,2020)
	want := uint64(1)
	if got := g.values[len(g.values)-1]; want != got {
		t.Error("TestFirst(",input,") ", got, " want ", want)
	}
}

func TestRun3(t *testing.T) {
	input := "2,1,3"
	g := LoadMap(input)
	g = RunGame(g,2020)
	want := uint64(10)
	if got := g.values[len(g.values)-1]; want != got {
		t.Error("TestFirst(",input,") ", got, " want ", want)
	}
}


func TestRun4(t *testing.T) {
	input := "1,2,3"
	g := LoadMap(input)
	g = RunGame(g,2020)
	want := uint64(27)
	if got := g.values[len(g.values)-1]; want != got {
		t.Error("TestFirst(",input,") ", got, " want ", want)
	}
}


func TestRun5(t *testing.T) {
	input := "2,3,1"
	g := LoadMap(input)
	g = RunGame(g,2020)
	want := uint64(78)
	if got := g.values[len(g.values)-1]; want != got {
		t.Error("TestFirst(",input,") ", got, " want ", want)
	}
}


func TestRun6(t *testing.T) {
	input := "3,2,1"
	g := LoadMap(input)
	g = RunGame(g,2020)
	want := uint64(438)
	if got := g.values[len(g.values)-1]; want != got {
		t.Error("TestFirst(",input,") ", got, " want ", want)
	}
}

func TestRun7(t *testing.T) {
	input := "3,1,2"
	g := LoadMap(input)
	g = RunGame(g,2020)
	want := uint64(1836)
	if got := g.values[len(g.values)-1]; want != got {
		t.Error("TestFirst(",input,") ", got, " want ", want)
	}
}

func TestRunDay2(t *testing.T) {
	inputs := []string{"0,3,6","1,3,2","2,1,3","1,2,3","2,3,1","3,2,1","3,1,2"}
	wants := []uint64{175594,2578,3544142,261214,6895259,18,362}
	length := uint64(30000000)
	for index, input := range inputs {
		g := LoadMap(input)
		g = RunGame(g,length)
		want := wants[index]
		if got := g.values[len(g.values)-1]; want != got {
			t.Error("TestFirst(",input,") ", got, " want ", want)
		}
	}
}