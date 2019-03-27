package main

import (
	"fmt"
	"sort"
)

type person struct {
	firstname string
	lastname  string
}

type personList []person

func (list personList) Len() int {
	return len(list)
}

func (list personList) Less(i, j int) bool {
	return list[i].lastname < list[j].lastname
}

func (list personList) Swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func main() {
	list := personList{
		{firstname: "Duong", lastname: "Pham"},
		{firstname: "An", lastname: "Doan"},
		{firstname: "Giang", lastname: "Nguyen"},
	}
	fmt.Println(list)
	sort.Sort(list)
	fmt.Println(list)
}
