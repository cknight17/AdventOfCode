package main

import (
	"testing"
	//"reflect"
	"fmt"
)

func TestReadFiles(t *testing.T) {
	tree := BuildTree(ReadFile("test.txt"))
	fmt.Println(tree)
	list := ListContainers(FindContainers("shiny gold",tree,make(map[string]bool,0)))
	fmt.Println(list,len(list))
}

func TestReadFilesProd(t *testing.T) {
	tree := BuildTree(ReadFile("prod.txt"))
	fmt.Println(tree)
	list := ListContainers(FindContainers("shiny gold",tree,make(map[string]bool,0)))
	fmt.Println(list,len(list))
}

func TestReadFilesTest2(t *testing.T) {
	tree := BuildReverseTree(ReadFile("test.txt"))
	fmt.Println(tree)
	list := FindContents("shiny gold",tree)
	fmt.Println(list,len(list))
}

func TestReadFilesProd2(t *testing.T) {
	tree := BuildReverseTree(ReadFile("prod.txt"))
	fmt.Println(tree)
	list := FindContents("shiny gold",tree)
	fmt.Println(list,len(list))
}