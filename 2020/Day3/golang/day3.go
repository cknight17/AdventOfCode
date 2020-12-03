package main

import (
	"aoc.knighttechnology.net/trajectory"
	"fmt"
)

func main() { 
	lines := trajectory.GetInput("prod.txt")
	x,y, count := trajectory.MoveAll(0,0,3,1,lines) 
	fmt.Println(x,y,count)

	input := [][]int{[]int{0,0,1,1},[]int{0,0,3,1},[]int{0,0,5,1},[]int{0,0,7,1},[]int{0,0,1,2}}
	got := 1
	for _, item := range input {
		x, y, count := trajectory.MoveAll(item[0],item[1],item[2],item[3],lines)
		got = got * count
		fmt.Println(x,y,count)
	}
	fmt.Println(got)
	// 7637517459 too high
	// 7560370818
} 