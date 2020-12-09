package main

import (
	"testing"
	"fmt"
	"sort"
)

func TestGetReindeer(t *testing.T) {
	input := "Vixen can fly 8 km/s for 8 seconds, but then must rest for 53 seconds."
	_, got := GetReindeer(input)
	want := true
	//fmt.Println(r)
	if want != got {
		t.Errorf("GetReindeer(%q) %t, want %t",input,got,want)
	}
}

func TestMovement(t *testing.T) {
	inputs := ReadFile("test.txt")
	reindeer := make([]Reindeer,0)
	for _, input := range inputs {
		r, ok := GetReindeer(input)
		if ok {
			reindeer = append(reindeer,r)
		} else {
			t.Errorf("GetReindeer(%q) %t %t",input,ok,true)
		}
	}
	ts := make([]TimeStamp,0)
	for _, current := range reindeer {
		ts = append(ts,ProgressAt(current,1000))
	}
	sort.Sort(ByDistance(ts))
	want := 1120
	max := ts[len(ts)-1]
	if got := max.distance; want != got {
		t.Errorf("GetReindeer(test.txt) %d, want %d\n",got,want)
		t.Error(t)
	}
}

func TestRace(t *testing.T) {
	inputs := ReadFile("test.txt")
	reindeer := make([]Reindeer,0)
	for _, input := range inputs {
		r, ok := GetReindeer(input)
		if ok {
			reindeer = append(reindeer,r)
		} else {
			t.Errorf("GetReindeer(%q) %t %t",input,ok,true)
		}
	}
	points := make(map[string]int)
	for _, r := range reindeer {
		points[r.name] = 0
	}

	for i := 1; i <= 1000; i++ {
		ts := make([]TimeStamp,0)
		for _, current := range reindeer {
			ts = append(ts,ProgressAt(current,i))
		}
		sort.Sort(ByDistance(ts))
		//fmt.Println(ts)
		max := ts[len(ts)-1]
		leadTime := max.distance
		for _, t := range ts {
			if t.distance == leadTime {
				points[t.reindeer.name]++
				fmt.Println("WINNER: ",i,t)
			} else {
				fmt.Println("LOSER: ",i,t)
			}
		}
	}
	fmt.Println(points)
}