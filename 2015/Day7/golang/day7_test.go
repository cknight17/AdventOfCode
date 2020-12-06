package main

import (
	"testing"
	"reflect"
)

func TestGetValue(t *testing.T) {
	instructions := Instructions(ReadFile("test.txt"))
	cachedValues = make(map[string]uint16)
	keys := []string{"d","e","f","g","h","i","x","y"}
	want := []uint16{72,507,492,114,65412,65079,123,456}
	got := make([]uint16,0)

	for _, key := range keys {
		got = append(got,GetValue(key, instructions))
	}
	
	if !reflect.DeepEqual(want,got) {
		t.Error("GetValue(test.txt)",got,want)
	}
}

