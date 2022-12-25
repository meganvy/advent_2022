// boilerplate
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {
	vals := []int{20, 60, 100, 140, 180, 220}
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 1
	x := 1
	sigStrength := 0
	for scanner.Scan() {
		l := scanner.Text()

		for _, val := range vals {
			if count == val {
				sigStrength += (count * x)
				break
			}
		}

		if l == "noop" {
			count += 1
		} else {
			line := strings.Split(l, " ")
			v, _ := strconv.Atoi(line[1])
			for _, val := range vals {
				if count < val && count+2 > val {
					sigStrength += (val * x)
				}
			}
			x += v
			count += 2
		}

		fmt.Printf("x is %d\t count is %d\n", x, count)

	}

	fmt.Println(sigStrength)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

}

// func part2() {
// 	vals := []int{20, 60, 100, 140, 180, 220}
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("failed to read input file")
// 		os.Exit(5)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	instructions := make(map[int]int)
// 	count := 1
// 	x := 1
// 	sigStrength := 0
// 	for scanner.Scan() {
// 		l := scanner.Text()
// 		if l != "noop" {
// 			line := strings.Split(l, " ")
// 			v, _ := strconv.Atoi(line[1])
// 			instructions[count+2] = v
// 		}
// 		count += 1
// 	}
// 	fmt.Println(instructions)

// 	for i := 1; i < 221; i++ {
// 		if v, ok := instructions[i]; ok {
// 			x += v
// 		}

// 		for _, val := range vals {
// 			if i == val {
// 				sigStrength += (val * x)
// 			}
// 		}
// 	}

// 	fmt.Println(sigStrength)

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(err)
// 		os.Exit(10)
// 	}

// }

func part2() {
	vals := []int{20, 60, 100, 140, 180, 220}
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 1
	x := 1
	sigStrength := 0
	for scanner.Scan() {
		l := scanner.Text()

		// for _, val := range vals {
		// 	if count == val {
		// 		sigStrength += (count * x)
		// 		break
		// 	}
		// }

		if l == "noop" {
			if math.Abs(float64(count)-float64(x)) < 2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			for _, val := range vals {
				if count == val {
					fmt.Print("\n")
				}
			}
			count += 1

		} else {
			line := strings.Split(l, " ")
			v, _ := strconv.Atoi(line[1])
			for _, val := range vals {
				if count < val && count+2 > val {
					fmt.Print("\n")
				}
			}
			if math.Abs(float64(count)-float64(x)) < 2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			x += v
			count += 2
		}

		// fmt.Printf("x is %d\t count is %d\n", x, count)

	}

	fmt.Println(sigStrength)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

}
