package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	//"strconv"
	//"sort"
)

type Coordinate struct {
	x int
	y int
}

func main() {
	mapl := LoadMap(ReadFile("prod.txt"))
	got := make([]int,0)
	for i := 0; i < 100; i++ {
		got = append(got, CountOn(mapl))
		//fmt.Println(mapl)
		mapl = NextFrame(mapl)
	}
	fmt.Println(got)
	fmt.Println(got[len(got)-1])
	fmt.Println(CountOn(mapl))
	// 799 too low
	mapl = LoadMap(ReadFile("prod.txt"))
	got = make([]int,0)
	for i := 0; i < 100; i++ {
		got = append(got, CountOn(mapl))
		//fmt.Println(mapl)
		mapl = NextFrame2(mapl)
	}
	fmt.Println(got)
	fmt.Println(got[len(got)-1])
	fmt.Println(CountOn(mapl))
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func LoadMap(inputs []string) [][]bool {
	mapl := make([][]bool,0)
	for _, input := range inputs {
		line := make([]bool,0)
		for _, char := range input {
			line = append(line,char == '#')
		}
		mapl = append(mapl,line)
	}
	return mapl
}

func GetNeighbors(at Coordinate, mapl [][]bool) []Coordinate {
	clist := make([]Coordinate,0)
	maxx := len(mapl)
	maxy := len(mapl[0])
	//fmt.Println(maxx,maxy)
	var c Coordinate
	// -1,-1
	for cx := at.x - 1; cx <= at.x + 1; cx++ {
		for cy := at.y - 1; cy <= at.y + 1; cy++ {
			//fmt.Println(cx,cy,at.x,at.y)
			if cx >= 0 && cx < maxx && cy >= 0 && cy < maxy && !(cx == at.x && cy == at.y) {
				c = Coordinate {
					x: cx,
					y: cy,
				}
				clist = append(clist,c)
			}
		}
	}
	return clist
}

func CountNeighbors(neighbors []Coordinate, mapl [][]bool) int {
	count := 0
	for _, coordinate := range neighbors {
		if mapl[coordinate.x][coordinate.y] {
			count++
		}
	}
	return count
}

func NextFrame(mapl [][]bool) [][]bool {
	nmapl := make([][]bool,0)
	for i, row := range mapl {
		ncol := make([]bool,0)
		for j, col := range row {
			currentState := col
			count := CountNeighbors(GetNeighbors(Coordinate { x: i, y: j },mapl),mapl)
			newState := false
			if currentState {
				newState = count == 2 || count == 3
			} else {
				newState = count == 3
			}
			ncol = append(ncol,newState)
		}
		nmapl = append(nmapl,ncol)
	}
	return nmapl
}

func NextFrame2(mapl [][]bool) [][]bool {
	nmapl := make([][]bool,0)
	maxx := len(mapl) - 1
	maxy := len(mapl[0]) - 1
	mapl[0][0] = true
	mapl[0][maxy] = true
	mapl[maxx][0] = true
	mapl[maxx][maxy] = true
	for i, row := range mapl {
		ncol := make([]bool,0)
		for j, col := range row {
			currentState := col
			count := CountNeighbors(GetNeighbors(Coordinate { x: i, y: j },mapl),mapl)
			newState := false
			if currentState {
				newState = count == 2 || count == 3
			} else {
				newState = count == 3
			}
			ncol = append(ncol,newState)
		}
		nmapl = append(nmapl,ncol)
	}
	nmapl[0][0] = true
	nmapl[0][maxy] = true
	nmapl[maxx][0] = true
	nmapl[maxx][maxy] = true
	return nmapl
}

func CountOn(mapl [][]bool) int {
	counter := 0
	for _, row := range mapl {
		for _, col := range row {
			if col {
				counter++
			}
		}
	}
	return counter
}