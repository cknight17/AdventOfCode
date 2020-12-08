package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	locations, routes := LoadPaths(ReadFile("prod.txt"))
	path, min := MinPath(locations,routes)
	fmt.Println(path,min)
	mpath, max := MaxPath(locations,routes)
	fmt.Println(mpath,max)
	// 464 too high
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func MinPath(locations []string, routes map[string]map[string]int) ([]string, int) {
	allPaths := AllPaths(locations)
	minPath := 0
	goldenPath := make([]string,0)
	minFound := false
	for _, path := range allPaths {
		//fmt.Println("PATH ",path)
		length := 0
		pathFound := true
		for i := 0; i < len(path)-1; i++ {
			from := path[i]
			to := path[i+1]
			route, ok := routes[from]
			if !ok {
				pathFound = false
				//fmt.Println("NOT FOUND from ",from,path)
				break
			} else {
				hop, ok := route[to]
				if !ok {
					//fmt.Println("NOT FOUND from ",from," to ", to,path)
					pathFound = false
					break
				}
				length += hop
			}
		}
		if pathFound {
			//fmt.Println("IS FOUND: ",length, path)
			if !minFound {
				//fmt.Println("STARTING ",length,path)
				minFound = true
				minPath = length
				goldenPath = path
			} else {
				if length < minPath {
					//fmt.Println("NEW MIN ",length, path)
					minFound = pathFound
					minPath = length
					goldenPath = path
				}
			}
		}
	}
	return goldenPath,minPath
}

func MaxPath(locations []string, routes map[string]map[string]int) ([]string, int) {
	allPaths := AllPaths(locations)
	minPath := 0
	goldenPath := make([]string,0)
	minFound := false
	for _, path := range allPaths {
		//fmt.Println("PATH ",path)
		length := 0
		pathFound := true
		for i := 0; i < len(path)-1; i++ {
			from := path[i]
			to := path[i+1]
			route, ok := routes[from]
			if !ok {
				pathFound = false
				//fmt.Println("NOT FOUND from ",from,path)
				break
			} else {
				hop, ok := route[to]
				if !ok {
					//fmt.Println("NOT FOUND from ",from," to ", to,path)
					pathFound = false
					break
				}
				length += hop
			}
		}
		if pathFound {
			//fmt.Println("IS FOUND: ",length, path)
			if !minFound {
				//fmt.Println("STARTING ",length,path)
				minFound = true
				minPath = length
				goldenPath = path
			} else {
				if length > minPath {
					//fmt.Println("NEW MIN ",length, path)
					minFound = pathFound
					minPath = length
					goldenPath = path
				}
			}
		}
	}
	return goldenPath,minPath
}

func LoadPaths(lines []string) ([]string,map[string]map[string]int) {
	pathFormat := "^(.+) to (.+) = ([0-9]+)$"
	r, _ := regexp.Compile(pathFormat)
	locations := make([]string,0)
	locationMap := make(map[string]bool,0)
	hops := make(map[string]map[string]int,0)
	for _, line := range lines {
		tokens := r.FindStringSubmatch(line)
		if len(tokens) == 4 {
			locationMap[tokens[1]] = true
			locationMap[tokens[2]] = true
			value, _ := strconv.ParseInt(tokens[3],10,32)
			hopSet, ok := hops[tokens[1]]
			if ok {
				hopSet[tokens[2]] = int(value)
			} else {
				hop := make(map[string]int)
				hop[tokens[2]] = int(value)
				hops[tokens[1]] = hop
			}
			hopSet2, ok2 := hops[tokens[2]]
			if ok2 {
				hopSet2[tokens[1]] = int(value)
			} else {
				hop := make(map[string]int)
				hop[tokens[1]] = int(value)
				hops[tokens[2]] = hop
			}
		} else {
			fmt.Println("ERROR ",tokens, " ",line)
		}
	}
	for key, _ := range locationMap {
		locations = append(locations,key)
	}
	return locations, hops
}

func Find(a []string, x string) (int, bool) {
	for i, n := range a {
		if x == n {
			return i, true
		}
	}
	return 0, false
}

func AllPaths(locations []string) [][]string {
	return permutations(locations)
}

func permutations(arr []string)[][]string{
    var helper func([]string, int)
    res := [][]string{}

    helper = func(arr []string, n int){
        if n == 1{
            tmp := make([]string, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}
