package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type position struct {
	y, x int
}

func (pos *position) updatePos(chr string) {
	switch chr {
	case "^":
		pos.y++
	case "v":
		pos.y--
	case ">":
		pos.x++
	case "<":
		pos.x--
	}
}

func contains(houses *[]position, house position) bool {
	for _, h := range *houses {
		if h.x == house.x && h.y == house.y {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	posSanta := position{0, 0}
	posRobot := position{0, 0}
	visitedHouses := map[position]bool{posSanta: true}
	robot := false

	for scanner.Scan() {
		var pos *position
		if robot {
			pos = &posRobot
		} else {
			pos = &posSanta
		}

		pos.updatePos(scanner.Text())
		visitedHouses[*pos] = true
		robot = !robot
	}
	fmt.Println(len(visitedHouses))
	checkError(scanner.Err())
}
