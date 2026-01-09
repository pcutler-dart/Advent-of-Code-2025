package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// support function to count the cl
func countZeroClicks(prev int, dir byte, dist int) int {
	if dist <= 0 {
		return 0
	}

	var offset int
	switch dir {
	case 'R':
		offset = (100 - prev) % 100
	case 'L':
		offset = prev % 100
	default:
		panic("invalid direction")
	}

	first := offset
	if first == 0 {
		first = 100 // starting on 0 means hitting 0 again after 100 clicks
	}

	if dist < first {
		return 0
	}
	return 1 + (dist-first)/100
}

func part1(lines []string) int {
	pos := 50
	countZero := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		dir := line[0] // 'L' or 'R'
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		dist = dist % 100 // rotating 100 clicks changes nothing

		switch dir {
		case 'R':
			pos = (pos + dist) % 100
		case 'L':
			// keep positive before mod
			pos = (pos - dist + 100) % 100
		default:
			panic("invalid direction: " + string(dir))
		}

		if pos == 0 {
			countZero++
		}
	}

	return countZero
}

// Part 2 assumes landing on 0 is not
func part2(lines []string) int {
	pos := 50
	countZero := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		prev := pos
		countZero += countZeroClicks(prev, dir, dist)

		// update final position
		rem := dist % 100
		switch dir {
		case 'R':
			pos = (prev + rem) % 100
		case 'L':
			pos = (prev - rem + 100) % 100
		}
	}

	return countZero
}

func main() {
	input := "input.txt"

	lines, err := readLines(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
