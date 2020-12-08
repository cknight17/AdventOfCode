package main

import "fmt"

func main() {
	input := "hxbxwxba"
	output, ok := NextValid(input)
	if ok {
		fmt.Println(output)
	} else {
		fmt.Println("ERROR")
	}
	output2, ok2 := NextValid(output)
	if ok2 {
		fmt.Println(output2)
	} else {
		fmt.Println("ERROR")
	}
}

func NextInput(input string) string {
	runes := []rune(input)
	if (runes[7] == 'z') {
		runes[7] = 'a'
		if (runes[6] == 'z') {
			runes[6] = 'a'
			if (runes[5] == 'z') {
				runes[5] = 'a'
				if (runes[4] == 'z') {
					runes[4] = 'a'
					if (runes[3] == 'z') {
						runes[3] = 'a'
						if (runes[2] == 'z') {
							runes[2] = 'a'
							if (runes[1] == 'z') {
								runes[1] = 'a'
								if (runes[0] == 'z') {
									return "zzzzzzzz"
								} else {
									runes[0]++
								}
							} else {
								runes[1]++
							}
						} else {
							runes[2]++
						}
					} else {
						runes[3]++
					}
				} else {
					runes[4]++
				}
			} else {
				runes[5]++
			}
		} else {
			runes[6]++
		}
	} else {
		runes[7]++
	}
	return string(runes)
}

func AllInputs(input string) []string {
	allInputs := make([]string,1)
	allInputs[0] = input
	runes := []rune(input)
	for string(runes) != "zzzzzzzz" {
		if (runes[7] == 'z') {
			runes[7] = 'a'
			if (runes[6] == 'z') {
				runes[6] = 'a'
				if (runes[5] == 'z') {
					runes[5] = 'a'
					if (runes[4] == 'z') {
						runes[4] = 'a'
						if (runes[3] == 'z') {
							runes[3] = 'a'
							if (runes[2] == 'z') {
								runes[2] = 'a'
								if (runes[1] == 'z') {
									runes[1] = 'a'
									if (runes[0] == 'z') {
										break
									} else {
										runes[0]++
									}
								} else {
									runes[1]++
								}
							} else {
								runes[2]++
							}
						} else {
							runes[3]++
						}
					} else {
						runes[4]++
					}
				} else {
					runes[5]++
				}
			} else {
				runes[6]++
			}
		} else {
			runes[7]++
		}
		allInputs = append(allInputs,string(runes))
	}
	return allInputs
}

func CheckIncreasing(input string) bool {
	for i := 0; i < len(input) - 2; i++ {
		rune1 := input[i]
		rune2 := input[i+1]
		rune3 := input[i+2]
		if (rune2 == rune1 + 1 && rune3 == rune2 + 1) {
			return true
		}
	}
	return false
}

func CheckForbidden(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == 'i' || input[i] == 'o' || input[i] == 'l' {
			return false
		}
	}
	return true
}

func CheckPairs(input string) bool {
	for i := 0; i < len(input) - 3; i++ {
		if input[i] == input[i+1] {
			for j := i+2; j < len(input) - 1; j++ {
				if input[j] == input[j+1] {
					return true
				}
			}
		}
	}
	return false
}

func CheckInput(input string) bool {
	return CheckIncreasing(input) && CheckForbidden(input) && CheckPairs(input)
}

func NextValid(input string) (string, bool) {
	for {
		if input == "zzzzzzzz" {
			return "zzzzzzzz", false
		}
		input = NextInput(input)
		if CheckInput(input) {
			return input, true
		}
	}
}