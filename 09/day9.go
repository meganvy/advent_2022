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
	part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}

	scanner := bufio.NewScanner(file)
	tail_visited := map[string]int{"0,0": 1}
	head_visited := map[string]int{"0,0": 1}
	hx, hy := 0, 0
	tx, ty := 0, 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		num, _ := strconv.Atoi(line[1])
		for i := 0; i < num; i++ {
			if dir == "U" {
				hy += 1
			} else if dir == "R" {
				hx += 1
			} else if dir == "D" {
				hy -= 1
			} else {
				hx -= 1
			}

			x_diff, y_diff := math.Abs(float64(hx)-float64(tx)), math.Abs(float64(hy)-float64(ty))

			if x_diff > 1 {
				if hx > tx {
					tx += 1
				} else {
					tx -= 1
				}
				if y_diff > 0 {
					ty = hy
				}
			} else if y_diff > 1 {
				if hy > ty {
					ty += 1
				} else {
					ty -= 1
				}
				if x_diff > 0 {
					tx = hx
				}
			}

			pos := fmt.Sprintf("%d,%d", tx, ty)
			tail_visited[pos] = 1
			fmt.Println(pos)
			head_visited[fmt.Sprintf("%d,%d", hx, hy)] = 1
		}
	}

	// fmt.Println(tail_visited)
	fmt.Println(len(tail_visited))

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
	visited := map[string]map[string]int{
		"1":  {"0,0": 1},
		"2":  {"0,0": 1},
		"3":  {"0,0": 1},
		"4":  {"0,0": 1},
		"5":  {"0,0": 1},
		"6":  {"0,0": 1},
		"7":  {"0,0": 1},
		"8":  {"0,0": 1},
		"9":  {"0,0": 1},
		"10": {"0,0": 1},
	}
	// x, y
	pos := map[int][]int{
		1:  {0, 0},
		2:  {0, 0},
		3:  {0, 0},
		4:  {0, 0},
		5:  {0, 0},
		6:  {0, 0},
		7:  {0, 0},
		8:  {0, 0},
		9:  {0, 0},
		10: {0, 0},
	}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		num, _ := strconv.Atoi(line[1])
		for i := 0; i < num; i++ {
			if dir == "U" {
				pos[1][1] += 1
			} else if dir == "R" {
				pos[1][0] += 1
			} else if dir == "D" {
				pos[1][1] -= 1
			} else {
				pos[1][0] -= 1
			}

			// iterate each knot
			for k := 1; k <= 9; k++ {
				hx, tx := pos[k][0], pos[k+1][0]
				hy, ty := pos[k][1], pos[k+1][1]
				x_diff, y_diff := math.Abs(float64(hx)-float64(tx)), math.Abs(float64(hy)-float64(ty))

				if x_diff > 1 {
					if hx > tx {
						tx += 1
					} else {
						tx -= 1
					}
					if y_diff > 0 {
						if hy > ty {
							ty += 1
						} else {
							ty -= 1
						}
					}
				} else if y_diff > 1 {
					if hy > ty {
						ty += 1
					} else {
						ty -= 1
					}
					if x_diff > 0 {
						if hx > tx {
							tx += 1
						} else {
							tx -= 1
						}
					}
				}

				pos[k][0], pos[k+1][0] = hx, tx
				pos[k][1], pos[k+1][1] = hy, ty
				visited[fmt.Sprint(k)][fmt.Sprintf("%d,%d", pos[k][0], pos[k][1])] = 1
				visited[fmt.Sprint(k+1)][fmt.Sprintf("%d,%d", pos[k+1][0], pos[k+1][1])] = 1
			}
			if dir == "U" && num == 8 {
				fmt.Println(pos)
				fmt.Println(len(visited["10"]))

			}
		}
	}

	fmt.Println(len(visited["10"]))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

	defer file.Close()
}
