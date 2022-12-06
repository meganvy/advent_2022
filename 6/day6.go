// boilerplate
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(part1(false))
	fmt.Println(part1(true))
}

func part1(part2 bool) int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}
	defer file.Close()

	values := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count := 0
		start := 0
		end := 4
		if part2 {
			end = 14
		}
		dup := 0
		for {
			dup = 0
			for _, key := range line[start:end] {
				if _, ok := values[string(key)]; ok {
					dup += 1
					values[string(key)] += 1
				} else {
					values[string(key)] = 1
				}
			}
			// fmt.Println(values)
			// fmt.Println(dup)
			if dup == 0 {
				addend := 4
				if part2 {
					addend = 14
				}
				return count + addend
				break
			} else {
				values = map[string]int{}
				count += 1
				start += 1
				end += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

	return -1
}
