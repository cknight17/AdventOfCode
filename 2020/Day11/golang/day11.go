package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"reflect"
)

func main() {
	input := ReadFile("prod.txt")
	for _, line := range input {
		fmt.Println(line)
	}
	fmt.Println("")
	input2 := NextFrame(input)
	for !reflect.DeepEqual(input,input2) {
		for _, line := range input2 {
			fmt.Println(line)
		}
		fmt.Println("")
		input = input2
		input2 = NextFrame(input)
	}
	for _, line := range input2 {
		fmt.Println(line)
	}
	fmt.Println("")
	counter := 0
	for _, line := range input2 {
		for _, r := range line {
			if r == '#' {
				counter++
			}
		}
	}
	fmt.Println(counter)
	Day2()
}

func Day2() {
	input := ReadFile("prod.txt")
	for _, line := range input {
		fmt.Println(line)
	}
	fmt.Println("")
	
	input2 := NextFrame2(input)
	for !reflect.DeepEqual(input,input2) {
		for _, line := range input2 {
			fmt.Println(line)
		}
		fmt.Println("")
		input = input2
		input2 = NextFrame2(input)
	}
	for _, line := range input2 {
		fmt.Println(line)
	}
	fmt.Println("")
	counter := 0
	for _, line := range input2 {
		for _, r := range line {
			if r == '#' {
				counter++
			}
		}
	}
	fmt.Println(counter)
}

type Coordinate struct {
	x int
	y int
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func GetNeighbors(at Coordinate, mapl []string) []Coordinate {
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

func CountNeighbors(neighbors []Coordinate, mapl []string) int {
	count := 0
	for _, coordinate := range neighbors {
		//fmt.Printf("%q",mapl[coordinate.x][coordinate.y])
		if mapl[coordinate.x][coordinate.y] == '#' {
			count++
		}
	}
	return count
}

func NextFrame(mapl []string) []string {
	nmapl := make([]string,0)
	for i, row := range mapl {
		ncol := ""
		for j, col := range row {
			currentState := col
			count := CountNeighbors(GetNeighbors(Coordinate { x: i, y: j },mapl),mapl)
			newState := col
			//fmt.Println(currentState,count)
			if currentState == '#' && count >= 4 {
				newState = 'L'
			} else if currentState == 'L' && count == 0 {
				newState = '#'
			}
			ncol = ncol + string(newState)
		}
		nmapl = append(nmapl,ncol)
	}
	return nmapl
}

func GetDirections() []Coordinate {
	cset := make([]Coordinate,0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				cset = append(cset,Coordinate{x:i,y:j})
			}
		}
	}
	return cset
}

func Look(at Coordinate, input []string, max Coordinate, direction Coordinate) bool {
	at.x += direction.x
	at.y += direction.y
	if at.x < max.x && at.y < max.y && at.x >= 0 && at.y >= 0 {
		//fmt.Println(direction,at,string(input[at.x][at.y]))
		r := input[at.x][at.y]
		if r == '#' {
			return true
		} else {
			return Look(at,input,max,direction)
		}
	}
	return false
}

func LookAll(at Coordinate,input []string) int {
	max := Coordinate{x:len(input),y:len(input[0])}
	directions := GetDirections()
	counter := 0
	for _, direction := range directions {
		if Look(Coordinate{x:at.x,y:at.y},input,max,direction) {
			counter++
		}
	}
	return counter
}

func LookP(at Coordinate, input []string, max Coordinate, direction Coordinate) bool {
	at.x += direction.x
	at.y += direction.y
	if at.x < max.x && at.y < max.y && at.x >= 0 && at.y >= 0 {
		//fmt.Println(direction,at,string(input[at.x][at.y]))
		r := input[at.x][at.y]
		if r == '#' {
			return true
		} else if r == '.' {
			return LookP(at,input,max,direction)
		} else {
			return false
		}
	}
	return false
}


func LookAllP(at Coordinate,input []string) int {
	max := Coordinate{x:len(input),y:len(input[0])}
	directions := GetDirections()
	counter := 0
	for _, direction := range directions {
		if LookP(Coordinate{x:at.x,y:at.y},input,max,direction) {
			counter++
		}
	}
	return counter
}

func NextFrame2(mapl []string) []string {
	nmapl := make([]string,0)
	for i, row := range mapl {
		ncol := ""
		for j, col := range row {
			currentState := col
			//countL := LookAll(Coordinate { x: i, y: j },mapl)
			countP := LookAllP(Coordinate { x: i, y: j },mapl)
			newState := col
			//fmt.Println(currentState,count)
			if currentState == '#' && countP >= 5 {
				newState = 'L'
			} else if currentState == 'L' && countP == 0 {
				newState = '#'
			}
			ncol = ncol + string(newState)
		}
		nmapl = append(nmapl,ncol)
	}
	return nmapl
}

