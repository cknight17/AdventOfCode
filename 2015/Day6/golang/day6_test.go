package main

import (
	"testing"
	"reflect"
)

func TestCommand(t *testing.T) {
	inputs := []string{"turn on 0,0 through 999,999","toggle 0,0 through 999,0","turn off 499,499 through 500,500"}
	wantCommands := []string{"turn on","toggle","turn off"}
	wantX1s := []int{0,0,499} 
	wantY1s := []int{0,0,499}
	wantX2s := []int{999,999,500}
	wantY2s := []int{999,0,500}

	gotCommands := make([]string,0)
	gotX1s := make([]int,0)
	gotY1s := make([]int,0)
	gotX2s := make([]int,0)
	gotY2s := make([]int,0)

	for _, input := range inputs {
		command, x1, y1, x2, y2 := GetCommand(input)
		gotCommands = append(gotCommands,command)
		gotX1s = append(gotX1s,x1)
		gotY1s = append(gotY1s,y1)
		gotX2s = append(gotX2s,x2)
		gotY2s = append(gotY2s,y2)
	}

	if !(reflect.DeepEqual(wantCommands,gotCommands) && reflect.DeepEqual(wantX1s,gotX1s) && reflect.DeepEqual(wantX2s,gotX2s) && reflect.DeepEqual(wantY1s,gotY1s) && reflect.DeepEqual(wantY2s,gotY2s)) {
		t.Error("COMMANDS INVALID")
		t.Error(wantCommands,gotCommands)
		t.Error(wantX1s,gotX1s)
		t.Error(wantX2s,gotX2s)
		t.Error(wantY1s,gotY1s)
		t.Error(wantY2s,gotY2s)
	} 
}

func TestExecution(t *testing.T) {
	inputs := []string{"turn on 0,0 through 999,999","toggle 0,0 through 999,0","turn off 499,499 through 500,500"}
	want := []int{1000000,999000,998996}
	got := []int{}

	grid := GetGrid()
	for _, input := range inputs {
		command,x1,y1,x2,y2 := GetCommand(input)
		grid = ExecuteCommand(grid,command,x1,y1,x2,y2)
		got = append(got,CountGrid(grid))
	}
	if (!reflect.DeepEqual(want,got)) {
		t.Error("Count Mismatch",got,want)
	}
}

func TestExecution2(t *testing.T) {
	inputs := []string{"turn on 0,0 through 0,0","toggle 0,0 through 999,999"}
	want := []int64{1,2000001}
	got := []int64{}

	grid := GetGrid2()
	for _, input := range inputs {
		command,x1,y1,x2,y2 := GetCommand(input)
		grid = ExecuteCommand2(grid,command,x1,y1,x2,y2)
		got = append(got,CountGrid2(grid))
	}
	if (!reflect.DeepEqual(want,got)) {
		t.Error("Count Mismatch",got,want)
	}
}