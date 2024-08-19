package main

import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	checkErr(err)
	count, pos := 0, 0
	for idx, chr := range data {
		if string(chr) == ")" {
			count--
		} else {
			count++
		}
		if (pos == 0) && (count < 0) {
			pos = idx + 1
		}
	}
	fmt.Printf("count: %d\nInitial position: %d\n", count, pos)
}
