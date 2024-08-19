package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// checkErr logs and exits if an error is encountered
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// parseString converts a string in the format "NxNxN" to a slice of integers
func parseString(str string) []int {
	parsed_arr := strings.Split(str, "x")

	if len(parsed_arr) != 3 {
		checkErr(errors.New(fmt.Sprintf("Incorrect number of items, got %d", len(parsed_arr))))
	}

	val_arr := make([]int, 3)

	for idx, num := range parsed_arr {
		num, err := strconv.Atoi(num)
		checkErr(err)
		val_arr[idx] = num
	}

	return val_arr
}

// calcWrappingPaper calculates the total wrapping paper needed for a box
func calcWrappingPaper(arr []int) int {
	x, y, z := arr[0], arr[1], arr[2]
	lw, wh, hl := (x * y), (y * z), (z * x)
	return 2*lw + 2*wh + 2*hl + min(lw, wh, hl)
}

// calcRibbon calculates the total ribbon needed for a box
func calcRibbon(arr []int) int {
	slices.Sort(arr)
	x, y, z := arr[0], arr[1], arr[2]
	return 2*x + 2*y + x*y*z
}

func main() {
	data, err := os.Open("input.txt")
	checkErr(err)
	defer data.Close()

	paper, ribbon := 0, 0
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		arr := parseString(scanner.Text())
		paper += calcWrappingPaper(arr)
		ribbon += calcRibbon(arr)
	}

	checkErr(scanner.Err())
	fmt.Printf("Paper: %d\nRibbon: %d\n", paper, ribbon)
}
