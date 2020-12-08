package main

import (
	"testing"
	"fmt"
)

func TestReadInstructions(t *testing.T) {
	instructions := ReadInstructions(ReadFile("test.txt"))
	want := 9
	if got := len(instructions); want != got {
		t.Errorf("ReadInstructions(test.txt) %d, want %d",got,want)
	}
}

func TestProcessInstructions(t *testing.T) {
	instructions := ReadInstructions(ReadFile("test.txt"))
	program := CreateProgram(instructions)
	program = ProcessInstructions(program)
	want := 5
	if got := program.accumulator; want != got {
		t.Errorf("ProcessInstructions(test.txt) %d, want %d\n",got,want)
		t.Error(program)
	}
}

func TestProcessInstructions2(t *testing.T) {
	instructions := ReadInstructions(ReadFile("test.txt"))
	program := CreateProgram(instructions)
	programL := ProcessInstructions2(program)
	want := 8
	fmt.Println(programL)
	if got := programL.accumulator; want != got {
		t.Errorf("ProcessInstructions(test.txt) %d, want %d\n",got,want)
		t.Error(programL)
	}
}