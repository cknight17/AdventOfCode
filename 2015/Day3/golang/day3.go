package main

import (
	"strconv"
	"fmt"
	"io/ioutil"
)

func main() {
	input := ReadFile("prod.txt")
	fmt.Println(len(GetHouses(input)))
	fmt.Println(len(GetHousesWithRobot(input)))
}

func ReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return string(b)
}

func GetLabel(x int, y int) string {
	return strconv.Itoa(x) + "x" + strconv.Itoa(y)
}

func GetHouses(input string) map[string]int {
	houses := make(map[string]int,0)
	atX := 0
	atY := 0
	houses[GetLabel(atX,atY)] = 1
	for _, direction := range input {
		moveX := 0
		moveY := 0
		switch direction {
			case '>':
				moveX++
			case '<':
				moveX--
			case '^':
				moveY++
			case 'v':
				moveY--
			default:
				fmt.Println("ERROR")
		}
		atX += moveX
		atY += moveY
		label := GetLabel(atX,atY)
		value, ok := houses[label]
		if ok {
			houses[label] = value + 1
		} else {
			houses[label] = 1
		}
	}
	return houses
}

func GetHousesWithRobot(input string) map[string]int {
	houses := make(map[string]int,0)
	atX := 0
	atY := 0

	robotAtX := 0
	robotAtY := 0

	houses[GetLabel(atX,atY)] = 1
	for index, direction := range input {
		moveX := 0
		moveY := 0
		switch direction {
			case '>':
				moveX++
			case '<':
				moveX--
			case '^':
				moveY++
			case 'v':
				moveY--
			default:
				fmt.Println("ERROR")
		}
		label := ""
		moveSanta := true
		if index % 2 == 1 {
			moveSanta = false
		}
		if moveSanta {
			atX += moveX
			atY += moveY
			label = GetLabel(atX,atY)
		} else {
			robotAtX += moveX
			robotAtY += moveY
			label = GetLabel(robotAtX,robotAtY)
		}
		
		value, ok := houses[label]
		if ok {
			houses[label] = value + 1
		} else {
			houses[label] = 1
		}
	}
	return houses
}