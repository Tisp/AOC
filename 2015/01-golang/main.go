package main

import (
	"fmt"
	"os"
)

const (
	FLOOR_UP   = "("
	FLOOR_DOWN = ")"
	BASEMENT   = -1
)

func main() {
	SolutionPartOne()
	SolutionPartTwo()
}

func ReadSample(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func SolutionPartOne() {
	input := ReadSample("./sample-1.txt")

	current_floor := 0
	for _, dir := range input {
		switch string(dir) {
		case FLOOR_UP:
			current_floor += 1
		case FLOOR_DOWN:
			current_floor -= 1
		}
	}
	fmt.Println(current_floor)
}

func SolutionPartTwo() {
	input := ReadSample("./sample-2.txt")

	current_floor := 0
	for pos, dir := range input {
		switch string(dir) {
		case FLOOR_UP:
			current_floor += 1
		case FLOOR_DOWN:
			current_floor -= 1
		}
		if current_floor == BASEMENT {
			fmt.Println(pos + 1)
		}
	}
}
