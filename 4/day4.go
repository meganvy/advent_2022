// boilerplate
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		elves := strings.Split(line, ",")
		first := strings.Split(elves[0], "-")
		fStart, _ := strconv.Atoi(first[0])
		fEnd, _ := strconv.Atoi(first[1])
		second := strings.Split(elves[1], "-")
		sStart, _ := strconv.Atoi(second[0])
		sEnd, _ := strconv.Atoi(second[1])

		if fStart <= sStart && fEnd >= sEnd || sStart <= fStart && sEnd >= fEnd {
			count += 1
		}

	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

	defer file.Close()
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		elves := strings.Split(line, ",")
		first := strings.Split(elves[0], "-")
		fStart, _ := strconv.Atoi(first[0])
		fEnd, _ := strconv.Atoi(first[1])
		second := strings.Split(elves[1], "-")
		sStart, _ := strconv.Atoi(second[0])
		sEnd, _ := strconv.Atoi(second[1])

		// leave me alone HAHA
		if fStart <= sStart && fEnd >= sEnd || fStart <= sStart && fEnd >= sStart || sStart <= fStart && sEnd >= fEnd || sStart <= fStart && sEnd >= fStart {
			count += 1
		}
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

	defer file.Close()
}
