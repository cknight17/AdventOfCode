package main

import (
	"testing"
	"reflect"
)

func TestStuff(t *testing.T) {
	want := 0
	got := 0
	if !reflect.DeepEqual(want,got) {
		t.Error("stuff",got,want)
	}
}