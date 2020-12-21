package main

import (
	"testing"
	"reflect"
)

func TestDesignations(t *testing.T) {
	tt := ParseSections(ReadFile("test.txt"))
	tt = FindErrors(tt)
	want := int64(71)
	if got := TicketScanningErrorRate(tt); want != got {
		t.Errorf("TicketScanningErrorRate(test.txt) %d, want %d",got,want)
	}
	want2 := map[int64]bool{1:true,2:true,3:true}
	if got2 := tt.errorTicket; !reflect.DeepEqual(want2,got2) {
		t.Error("Error tickets: got ",got2," want ",want2)
	}
}

func TestDesignationLookup(t *testing.T) {
	tt := ParseSections(ReadFile("test2.txt"))
	tt = FindErrors(tt)
	tt = ApplyDesignations(tt)
	tt = ValidateDesignations(tt)
	want := map[string]int64{"class":12,"row":11,"seat":13}
	if got := tt.myFinalTicket; !reflect.DeepEqual(want,got) {
		t.Error("ValidateDesignations(test2.txt) ",got," want ",want)
	}
}