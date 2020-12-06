package main

import (
	"regexp"
	"strconv"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputs := ReadFile("prod.txt")
	grid := GetGrid()
	for _, input := range inputs {
		command,x1,y1,x2,y2 := GetCommand(input)
		grid = ExecuteCommand(grid,command,x1,y1,x2,y2)
	}
	fmt.Println(CountGrid(grid))
	grid2 := GetGrid2()
	for _, input := range inputs {
		command,x1,y1,x2,y2 := GetCommand(input)
		grid2 = ExecuteCommand2(grid2,command,x1,y1,x2,y2)
	}
	fmt.Println(CountGrid2(grid2))
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func GetGrid() [][]bool {
	grid := make([][]bool,0)
	for i := 0; i < 1000; i++ {
		row := make([]bool,0)
		for j := 0; j < 1000; j++ {
			row = append(row,false)
		}
		grid = append(grid,row)
	}
	return grid
}

func GetGrid2() [][]int {
	grid := make([][]int,0)
	for i := 0; i < 1000; i++ {
		row := make([]int,0)
		for j := 0; j < 1000; j++ {
			row = append(row,0)
		}
		grid = append(grid,row)
	}
	return grid
}

func CountGrid(input [][]bool) int {
	counter := 0
	for _, row := range input {
		for _, light := range row {
			if light {
				counter++
			}
		}
	}
	return counter
}

func CountGrid2(input [][]int) int64 {
	var counter int64 = 0
	for _, row := range input {
		for _, light := range row {
			counter += int64(light)
		}
	}
	return counter
}

func GetCommand(input string) (string,int,int,int,int) {
	commandString := "^(turn on|turn off|toggle) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)\n*$"
	r, _ := regexp.Compile(commandString)
	output := r.FindStringSubmatch(input)
	if len(output) < 6 {
		return "ERROR",-1,-1,-1,-1
	}
	command := output[1]
	var tmp int64
	tmp,_ = strconv.ParseInt(output[2],10,32)
	x1 := int(tmp)
	tmp,_ = strconv.ParseInt(output[3],10,32)
	y1 := int(tmp)
	tmp,_ = strconv.ParseInt(output[4],10,32)
	x2 := int(tmp)
	tmp,_ = strconv.ParseInt(output[5],10,32)
	y2 := int(tmp)
	if x1 > x2 {
		xtemp := x1
		x1 = x2
		x2 = xtemp
	}
	if y1 > y2 {
		ytemp := y1
		y1 = y2
		y1 = ytemp
	}
	return command,x1,y1,x2,y2
}

func ExecuteCommand(input [][]bool, command string, x1 int, y1 int, x2 int, y2 int) [][]bool {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			switch command {
				case "turn on":
					input[x][y] = true
				case "turn off":
					input[x][y] = false
				case "toggle":
					input[x][y] = !input[x][y]
			}
		}
	}
	return input
}

func ExecuteCommand2(input [][]int, command string, x1 int, y1 int, x2 int, y2 int) [][]int {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			switch command {
				case "turn on":
					input[x][y]++
				case "turn off":
					input[x][y]--
					if input[x][y] < 0 {
						input[x][y] = 0
					}
				case "toggle":
					input[x][y] += 2
			}
		}
	}
	return input
}