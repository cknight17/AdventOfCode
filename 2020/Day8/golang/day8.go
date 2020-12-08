package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	instructions := ReadInstructions(ReadFile("prod.txt"))
	program := CreateProgram(instructions)
	program = ProcessInstructions(program)
	fmt.Println(program)
	fmt.Println(program.accumulator)
	program = CreateProgram(instructions)
	program = ProcessInstructions2(program)
	fmt.Println(program)
	fmt.Println(program.accumulator)
}

type Instruction struct {
	command string
	value int
}

type Program struct {
	instructions []Instruction
	accumulator int
	at int
	visited map[int]bool
	complete bool
	fail bool
	path []int
	canChange bool
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func CreateProgram(instructions []Instruction) Program {
	var program Program
	program.instructions = instructions
	program.accumulator = 0
	program.at = 0
	program.visited = make(map[int]bool,0)
	program.complete = false
	program.fail = false
	program.path = make([]int,0)
	program.canChange = true
	return program
}

func CopyProgram(oldprogram Program) Program {
	var program Program
	program.instructions = oldprogram.instructions
	program.accumulator = oldprogram.accumulator
	program.at = oldprogram.at
	program.visited = make(map[int]bool,0)
	for key, value := range oldprogram.visited {
		program.visited[key] = value
	}
	program.complete = oldprogram.complete
	program.fail = oldprogram.fail
	program.path = make([]int,len(oldprogram.path))
	copy(program.path,oldprogram.path)
	program.canChange = oldprogram.canChange
	return program
}

func ReadInstructions(inputs []string) []Instruction {
	instructions := make([]Instruction,0)
	r := regexp.MustCompile(`^(nop|acc|jmp) ([\+\-][0-9]{1,3})$`)
	for _, input := range inputs {
		matches := r.FindAllStringSubmatch(input,-1)
		if len(matches) < 1 {
			fmt.Printf("ERROR no match for %q\n",input)
			fmt.Println(matches)
		} else {
			match := matches[0]
			if len(match) < 3 {
				fmt.Printf("ERROR no inner match for %q\n",input)
				fmt.Println(match)
			} else {
				var instruction Instruction
				instruction.command = match[1]
				value,_ := strconv.ParseInt(match[2],10,32)
				instruction.value = int(value)
				instructions = append(instructions,instruction)
			}

		}
	}
	//fmt.Println(instructions)
	return instructions
}

func ProcessInstructions(program Program) Program {
	if program.complete {
		return program
	}
	currentInstruction := program.instructions[program.at]
	program.visited[program.at] = true
	switch currentInstruction.command {
	case "nop":
		program.at++
		return ProcessInstructions(program)
	case "acc":
		program.accumulator += currentInstruction.value
		program.at++
		return ProcessInstructions(program)
	case "jmp":
		location := program.at + currentInstruction.value
		_, ok := program.visited[location]
		if ok {
			program.complete = true
			return program
		} else {
			program.at = location
			return ProcessInstructions(program)
		}
	}
	return program
}

func ProcessInstructions2(program Program) Program {
	if program.complete || program.fail {
		return program
	}
	if program.at == len(program.instructions)-1 {
		program.complete = true
	} else if program.at >= len(program.instructions) {
		program.fail = true
		return program
	}
	currentInstruction := program.instructions[program.at]
	program.visited[program.at] = true
	program.path = append(program.path,program.at)
	switch currentInstruction.command {
	case "nop":
		if program.canChange {
			programB := CopyProgram(program)
			programB = noop(programB)
			program.canChange = false
			program = jmp(program)
			if program.complete {
				return program
			} else {
				return programB
			}
		} else {
			return noop(program)
		}
	case "acc":
		program.accumulator += currentInstruction.value
		program.at++
		return ProcessInstructions2(program)
	case "jmp":
		if program.canChange {
			programB := CopyProgram(program)
			programB.canChange = false
			programB = noop(programB)
			program = jmp(program)
			if programB.complete {
				return programB
			} else {
				return program
			}
		} else {
			return jmp(program)
		}
	}
	return program
}

func noop(program Program) Program {
	program.at++
	return ProcessInstructions2(program)
}

func jmp(program Program) Program {
	location := program.at + program.instructions[program.at].value
	_, ok := program.visited[location]
	if ok {
		program.fail = true
		return program
	} else {
		program.at = location
		return ProcessInstructions2(program)
	}
}