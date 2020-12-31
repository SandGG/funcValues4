package main

import "fmt"

type data func(string) []string

func main() {
	print(fn)
}

func print(a data) {
	fmt.Println(a("Data"))
}

func fn(a string) []string {
	var writeNames = make([]string, 3)
	for i := range writeNames {
		writeNames[i] = a
	}
	return writeNames
}
