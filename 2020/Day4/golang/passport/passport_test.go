package passport

import (
	"testing"
	"reflect"
	//"fmt"
)

func TestReadFile(t *testing.T) {
	got := len(ReadFile("test.txt"))
	if got <= 0 {
		t.Errorf("ReadFile(test.txt) %d, want > 0",got)
	}
}

func TestGetPassports(t *testing.T) {
	want := 4
	item,_:=GetPassports(ReadFile("test.txt"))
	if got := len(item); want != got {
		t.Errorf("GetPassports(test.txt) %d, want > 0",got)
	}
}

func TestCheckRequiredFields(t *testing.T) {
	want := []bool{true,false,true,false}
	got := make([]bool,0)
	_,results := GetPassports(ReadFile("test.txt"))
	for _, item := range results {
		got = append(got,CheckRequiredFields(item))
	}
	//fmt.Println(got)
	if !reflect.DeepEqual(want,got) {
		t.Errorf("CheckRequiredFields(test.txt) %v, want %v",got,want)
	}
}

func TestNumberOfValidPassports(t *testing.T) {
	want := 2
	_,results := GetPassports(ReadFile("test.txt"))
	//fmt.Println(got)
	if got := NumberOfValidPassports(results); want!=got {
		t.Errorf("CheckRequiredFields(test.txt) %d, want %d",got,want)
	}
}

func TestCheckByr1(t *testing.T) {
	want := true
	input := "2002"
	if got := checkByr(input); want!=got {
		t.Errorf("checkByr(%q) %t, want %t",input,got,want)
	}
}

func TestCheckByr2(t *testing.T) {
	want := false
	input := "2003"
	if got := checkByr(input); want!=got {
		t.Errorf("checkByr(%q) %t, want %t",input,got,want)
	}
}

func TestCheckHgt(t *testing.T) {
	want := []bool{true,true,false,false}
	input := []string {"60in","190cm","190in","190"}
	got := make([]bool,0)
	for _, item := range input {
		got = append(got,checkHgt(item))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("checkHgt(%v) %v, want %v",input,got,want)
	}
}

func TestCheckHcl(t *testing.T) {
	want := []bool{true,true,false}
	input := []string {"#123abc","#123abz","123abc"}
	got := make([]bool,0)
	for _, item := range input {
		got = append(got,checkHcl(item))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("checkHcl(%v) %v, want %v",input,got,want)
	}
}

func TestCheckEcl(t *testing.T) {
	want := []bool{true,false}
	input := []string {"brn","wat"}
	got := make([]bool,0)
	for _, item := range input {
		got = append(got,checkEcl(item))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("checkEcl(%v) %v, want %v",input,got,want)
	}
}

func TestCheckPid(t *testing.T) {
	want := []bool{true,false}
	input := []string {"000000001","0123456789"}
	got := make([]bool,0)
	for _, item := range input {
		got = append(got,checkPid(item))
	}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("checkPid(%v) %v, want %v",input,got,want)
	}
}

func TestNumberOfValidVerifiedPassports(t *testing.T) {
	want := 0
	passports,results := GetPassports(ReadFile("testb1.txt"))
	//fmt.Println(got)
	if got := NumberOfValidVerifiedPassports(passports,results); want!=got {
		t.Errorf("NumberOfValidVerifiedPassports(testb1.txt) %d, want %d",got,want)
	}
}

func TestNumberOfValidVerifiedPassports2(t *testing.T) {
	want := 4
	passports,results := GetPassports(ReadFile("testb2.txt"))
	//fmt.Println(got)
	if got := NumberOfValidVerifiedPassports(passports,results); want!=got {
		t.Errorf("NumberOfValidVerifiedPassports(testb2.txt) %d, want %d",got,want)
	}
}