// boilerplate
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var stacks map[int][]string
var total int

func main() {
	fmt.Println(part1(true))
	fmt.Println(part2())
}

func part2() string {
	return part1(false)
}

func part1(part2 bool) string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}

	defer file.Close()

	stacks = make(map[int][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[") {
			count := 0
			for _, char := range line {
				if count%4 == 1 && unicode.IsLetter(char) {
					stacks[count] = append(stacks[count], string(char))
				}
				count += 1
			}
		} else if line == "" {
			// finished parsing crane stacks, reverse lists because prepending is 2 much
			for key, stack := range stacks {
				stacks[key] = reverse(stack)
			}
			fmt.Println(stacks, "initialized")

		} else if strings.HasPrefix(line, "move") {
			// hacc
			vals := strings.Split(line, " ")
			num, _ := strconv.Atoi(vals[1])
			s, _ := strconv.Atoi(vals[3])
			e, _ := strconv.Atoi(vals[5])
			start := 1 + (4 * (s - 1))
			end := 1 + (4 * (e - 1))

			// tried to move this to a helper but my debugging loglines suffered
			removedStack := stacks[start][len(stacks[start])-num:]
			newStartStack := stacks[start][:len(stacks[start])-num]
			stacks[start] = newStartStack
			useStack := reverse(removedStack)
			if part2 {
				useStack = removedStack
			}
			for _, char := range useStack {
				stacks[end] = append(stacks[end], char)
			}

		}

	}

	val := ""
	// lol at hardcoded 100 instead of iterating keys in sorted order
	for i := 0; i <= 100; i++ {
		if len(stacks[i]) != 0 {
			val = val + stacks[i][len(stacks[i])-1]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
	return val
}

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}
