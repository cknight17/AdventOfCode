package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

func main() {

}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func BuildTree(input []string) map[string][]string {
	removePeriod := regexp.MustCompile(`\.$`)
	container := regexp.MustCompile("^(.+) bags contain (.+)$")
	roots := make(map[string][]string,0)
	for _, item := range input {
		item = removePeriod.ReplaceAllString(item,"")
		bags := container.FindStringSubmatch(item)
		from := bags[1]
		tos := make([]string,0)
		if bags[2] != "no other bags" {
			containsBags := strings.Split(bags[2],",")
			for _, item := range containsBags {
				bagtype := regexp.MustCompile(`^\s*[0-9]+ (.+) (bag|bags)$`)
				//fmt.Println(item)
				bag := bagtype.FindStringSubmatch(item)
				
				tos = append(tos,bag[1])
			}
		}
		for _, item := range tos {
			root, ok := roots[item]
			if ok {
				root = append(root,from)
				roots[item] = root
			} else {
				root = make([]string,0)
				root = append(root,from)
				roots[item] = root
			}
		}
	}
	return roots
}

func BuildReverseTree(input []string) map[string]map[string]int {
	removePeriod := regexp.MustCompile(`\.$`)
	container := regexp.MustCompile("^(.+) bags contain (.+)$")
	roots := make(map[string]map[string]int,0)
	for _, item := range input {
		item = removePeriod.ReplaceAllString(item,"")
		bags := container.FindStringSubmatch(item)
		from := bags[1]
		tos := make(map[string]int,0)
		if bags[2] != "no other bags" {
			containsBags := strings.Split(bags[2],",")
			for _, item := range containsBags {
				bagtype := regexp.MustCompile(`^\s*([0-9]+) (.+) (bag|bags)$`)
				//fmt.Println(item)
				bag := bagtype.FindStringSubmatch(item)
				numBags, _ := strconv.Atoi(bag[1])
				tos[bag[2]] = int(numBags)
				roots[from] = tos
			}
		}
	}
	return roots
}

func FindContainers(input string, containers map[string][]string, current map[string]bool) map[string]bool {
	container, ok := containers[input]
	if ok {
		//fmt.Println(container)
		for _, item := range container {
			current[item] = true
			current = FindContainers(item,containers,current)
		}
	}
	return current
}

func FindContents(input string, contentsMap map[string]map[string]int) []string {
	contents := make([]string,0)
	items, ok := contentsMap[input]
	if ok {
		for key, value := range items {
			for i := 0; i < value; i++ {
				contents = append(contents,key)
				contents = append(contents,FindContents(key,contentsMap)...)
			}
		}
	}
	return contents
}

func ListContainers(input map[string]bool) []string {
	returnValue := make([]string,0)
	for item, _ := range input {
		returnValue = append(returnValue,item)
	}
	return returnValue
}