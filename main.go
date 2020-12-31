package main

import "fmt"

type search func(string) string

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(nameSearch(writeData())("Romina"))
}

func nameSearch(per []person) search {
	return func(str string) string {
		for _, v := range per {
			if str == v.name {
				return str + " found"
			}
		}
		return str + " not found"
	}
}

func writeData() []person {
	var people = []person{
		{name: "Maria", age: 13},
		{name: "Rosa", age: 29},
		{name: "Laura", age: 30},
		{name: "Juan", age: 25},
	}
	return people
}
