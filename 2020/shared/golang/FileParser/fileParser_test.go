package fileParser

import (
	"testing"
	"regexp"
	"reflect"
)

func TestLinereader(t *testing.T) {
	lines := Linereader("test.txt")
	want := 3
	if got := len(lines); want != got {
		t.Errorf("Linereader(test.txt) = %q, want %q",got,want)
	}
}

func TestDay2Regexp(t *testing.T) {
	day2LineFormat := "^([0-9]+)-([0-9]+) ([a-zA-Z]{1}): (.+)$"
	sampleLine := "1-3 a: abcde"
	r, _ := regexp.Compile(day2LineFormat)
	matches := r.FindStringSubmatch(sampleLine)
	want := []string{sampleLine,"1","3","a","abcde"}
	if !reflect.DeepEqual(want,matches) {
		t.Errorf("Hello() = %q, want %q",matches,want)
	}
}

func TestDay2parseline_scenario1(t *testing.T) {
	line := "1-3 a: abcde"
	want := Day2line{Min:1,Max:3,Character:"a",Input:"abcde"}
	got := day2parseline(line)
	if !reflect.DeepEqual(want,got) {
		t.Errorf("day2parseline(%q) = %q, want %q",line,got,want)
	}
}

func TestDay2parseline_scenario2(t *testing.T) {
	line := "1-3 b: cdefg"
	want := Day2line{Min:1,Max:3,Character:"b",Input:"cdefg"}
	got := day2parseline(line)
	if !reflect.DeepEqual(want,got) {
		t.Errorf("day2parseline(%q) = %q, want %q",line,got,want)
	}
}

func TestDay2parseline_scenario3(t *testing.T) {
	line := "2-9 c: ccccccccc"
	want := Day2line{Min:2,Max:9,Character:"c",Input:"ccccccccc"}
	got := day2parseline(line)
	if !reflect.DeepEqual(want,got) {
		t.Errorf("day2parseline(%q) = %q, want %q",line,got,want)
	}
}

func TestDay2input(t *testing.T) {
	file := "test.txt"
	want := 3
	output := Day2input(file)
	if got := len(output); want != got {
		t.Errorf("day2input(%q) = %d, want %d",file,got,want)
	}
}

func TestDay2input2(t *testing.T) {
	file := "test.txt"
	want := 3
	output := Day2input(file)
	if got := len(output); want != got {
		t.Errorf("day2input(%q) = %d, want %d",file,got,want)
	}
}