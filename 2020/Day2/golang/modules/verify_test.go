package verify

import (
	"testing"
	"aoc.knighttechnology.net/fileParser"
	"reflect"
)

func TestVerify1(t *testing.T) {
	want := true
	input := fileParser.Day2line{Min:1,Max:3,Character:"a",Input:"abcde"}
	if got := Verify(input); want != got {
	 	t.Errorf("verify(%q) = %t, want %t",input,got,want)
	}
}

func TestVerify2(t *testing.T) {
	want := true
	input := fileParser.Day2line{Min:2,Max:9,Character:"c",Input:"ccccccccc"}
	if got := Verify(input); want != got {
	 	t.Errorf("verify(%q) = %t, want %t",input,got,want)
	}
}

func TestVerify3(t *testing.T) {
	want := false
	input := fileParser.Day2line{Min:1,Max:3,Character:"b",Input:"cdefg"}
	if got := Verify(input); want != got {
	 	t.Errorf("verify(%q) = %t, want %t",input,got,want)
	}
}

func TestVerifyAll(t *testing.T) {
	want := []bool{true,true,false}
	file := "test.txt"
	output := fileParser.Day2input(file)
	if got:= VerifyAll(output); !reflect.DeepEqual(want,got) {
		t.Errorf("verifyAll(%q) = %t, want %t",file,got,want)
   }
}

func TestVerifyAllCount(t *testing.T) {
	want := 2
	file := "test.txt"
	output := fileParser.Day2input(file)
	if got:= VerifyAllCount(output); want != got {
		t.Errorf("verifyAll(%q) = %d, want %d",file,got,want)
   }
}

func TestVerify1b(t *testing.T) {
	want := true
	input := fileParser.Day2line{Min:1,Max:3,Character:"a",Input:"abcde"}
	if got := Verifyb(input); want != got {
	 	t.Errorf("verify(%q) = %t, want %t",input,got,want)
	}
}

func TestVerify2b(t *testing.T) {
	want := false
	input := fileParser.Day2line{Min:2,Max:9,Character:"c",Input:"ccccccccc"}
	if got := Verifyb(input); want != got {
	 	t.Errorf("verify(%q) = %t, want %t",input,got,want)
	}
}

func TestVerify3b(t *testing.T) {
	want := false
	input := fileParser.Day2line{Min:1,Max:3,Character:"b",Input:"cdefg"}
	if got := Verifyb(input); want != got {
	 	t.Errorf("verify(%q) = %t, want %t",input,got,want)
	}
}

func TestVerifyAllb(t *testing.T) {
	want := []bool{false,true,false}
	file := "test.txt"
	output := fileParser.Day2input(file)
	if got:= VerifyAllb(output); !reflect.DeepEqual(want,got) {
		t.Errorf("verifyAll(%q) = %t, want %t",file,got,want)
   }
}

func TestVerifyAllCountb(t *testing.T) {
	want := 1
	file := "test.txt"
	output := fileParser.Day2input(file)
	if got:= VerifyAllCountb(output); want != got {
		t.Errorf("verifyAll(%q) = %d, want %d",file,got,want)
   }
}