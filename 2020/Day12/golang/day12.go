package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
)

func main() {
	lines := ReadFile("prod.txt")
	commands := ParseCommands(lines)
	fmt.Println(commands)
	origin := Coordinate{x:0, y:0, d:'E'}
	destination := Navigate(origin,commands)
	fmt.Println(origin," => ",destination)
	distance := int(math.Abs(float64(destination.x)) + math.Abs(float64(destination.y)))
	fmt.Println(distance)

	origin = Coordinate{x:0, y:0, d:'E'}
	waypoint := Waypoint{a:[]WaypointAxis{WaypointAxis{p:10,d:'E'},WaypointAxis{p:1,d:'N'}}}
	destination = Navigate2(origin,waypoint,commands)
	fmt.Println(origin," => ",destination)
	distance = int(math.Abs(float64(destination.x)) + math.Abs(float64(destination.y)))
	fmt.Println(distance)
	// 44326 too high
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func ParseCommands(inputs []string) []Command {
	commands := make([]Command,0)
	for _, input := range inputs {
		commands = append(commands,ParseCommand(input))
	}
	return commands
}

type Coordinate struct {
	x int
	y int
	d rune
}

type WaypointAxis struct {
	p int
	d rune
}

type Waypoint struct {
	a []WaypointAxis
}

type Command struct {
	op rune
	value int
}

var directions []rune = []rune{'E','S','W','N'}

func Rotate(d rune,deg int) rune {
	direction := 0
	switch d {
		case 'E':
			direction = 0
		case 'S':
			direction = 1
		case 'W':
			direction = 2
		case 'N':
			direction = 3
	}
	modifier := deg / 90
	direction = (direction + modifier) % 4
	if direction < 0 {
		direction += 4
	}
	return directions[direction]
} 

func ParseCommand(input string) Command {
	op := input[0]
	val := input[1:]
	value, _ := strconv.ParseInt(val,10,32)
	return Command{ op:rune(op), value:int(value)}
}

func Navigate(origin Coordinate,commands []Command) Coordinate {
	at := Coordinate{x:origin.x, y:origin.y ,d:origin.d}
	for _, command := range commands {
		d := command.op
		switch command.op {
		case 'L':
			at.d = Rotate(at.d,command.value*-1)
			continue
		case 'R':
			at.d = Rotate(at.d,command.value)
			continue
		case 'F':
			d = at.d
		}
		switch d {
		case 'N':
			at.y = at.y - command.value
		case 'S':
			at.y = at.y + command.value
		case 'E':
			at.x = at.x + command.value
		case 'W':
			at.x = at.x - command.value
		case 'F':

		}
	}
	return at
}

func Navigate2(origin Coordinate,waypoint Waypoint,commands []Command) Coordinate {
	at := Coordinate{x:origin.x, y:origin.y ,d:origin.d}
	for _, command := range commands {
		//fmt.Println(string(command.op)," ",command.value)
		d := command.op
		switch command.op {
		case 'L':
			for index, wa := range waypoint.a {
				waypoint.a[index].d = Rotate(wa.d,command.value*-1)
			}
			continue
		case 'R':
			for index, wa := range waypoint.a {
				waypoint.a[index].d = Rotate(wa.d,command.value)
			}
			continue
		case 'F':
			for i := 0; i < command.value; i++ {
				for _, wa := range waypoint.a {
					at = Navigate(at,[]Command{Command{op:wa.d,value:wa.p}})
				}
			}
		}
		for index, wa := range waypoint.a {
			waypoint.a[index] = ProcessAxis(wa,d,command.value)
		}
	}
	return at
}

func ProcessAxis(input WaypointAxis,d rune, value int) WaypointAxis {
	if GetCardinal(input.d) != GetCardinal(d) || GetCardinal(d) == "X" {
		return input
	}
	multiplier := 1
	if input.d != d {
		multiplier = - 1
	}
	v := input.p + (value*multiplier)
	nd := input.d
	if v < 0 {
		v = v * -1
		fmt.Println("FLIP")
		nd = d
	}
	return WaypointAxis{p:v,d:nd}
}

func GetCardinal(d rune) string {
	if d == 'W' || d == 'E' {
		return "EW"
	} else if d == 'N' || d == 'S' { 
		return "NS"
	} else {
		return "X"
	}
}