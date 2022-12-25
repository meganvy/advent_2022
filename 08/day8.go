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
	part1()
	part2()
}

// don't @ me
var treeMap map[int][]string
var betterTreeMap map[int][]int
var maxTop map[string]int
var maxBottom map[string]int
var maxRight map[string]int
var maxLeft map[string]int
var count int

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}
	defer file.Close()

	treeMap = make(map[int][]string)
	betterTreeMap = make(map[int][]int)
	maxTop = make(map[string]int)
	maxBottom = make(map[string]int)
	maxRight = make(map[string]int)
	maxLeft = make(map[string]int)
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		l := scanner.Text()
		line := strings.Split(l, "")
		treeMap[count] = line
		count += 1
	}

	// top/left values
	for i := 0; i < count; i++ {
		betterTreeMap[i] = make([]int, count)
		for j := 0; j < len(treeMap[i]); j++ {
			indexString := fmt.Sprintf("%d,%d", i, j)
			previIndexString := fmt.Sprintf("%d,%d", i-1, j)
			prevjIndexString := fmt.Sprintf("%d,%d", i, j-1)
			val, _ := strconv.Atoi(fmt.Sprintf("%s", treeMap[i][j]))

			betterTreeMap[i][j] = val

			if i == 0 {
				maxTop[indexString] = val
			} else {
				if prevMax, ok := maxTop[previIndexString]; ok {
					maxTop[indexString] = int(math.Max(float64(prevMax), float64(val)))
				}
			}

			if j == 0 {
				maxLeft[indexString] = val
			} else {
				if prevMax, ok := maxLeft[prevjIndexString]; ok {
					maxLeft[indexString] = int(math.Max(float64(prevMax), float64(val)))
				}
			}
		}
	}

	// bottom/right values
	for i := len(treeMap) - 1; i >= 0; i-- {
		for j := len(treeMap[i]) - 1; j >= 0; j-- {
			indexString := fmt.Sprintf("%d,%d", i, j)
			previIndexString := fmt.Sprintf("%d,%d", i+1, j)
			prevjIndexString := fmt.Sprintf("%d,%d", i, j+1)
			val, _ := strconv.Atoi(fmt.Sprintf("%s", treeMap[i][j]))
			if i == len(treeMap)-1 {
				maxBottom[indexString] = val
			} else {
				if prevMax, ok := maxBottom[previIndexString]; ok {
					maxBottom[indexString] = int(math.Max(float64(prevMax), float64(val)))
				}
			}

			if j == len(treeMap[i])-1 {
				maxRight[indexString] = val
			} else {
				if prevMax, ok := maxRight[prevjIndexString]; ok {
					maxRight[indexString] = int(math.Max(float64(prevMax), float64(val)))
				}
			}
		}
	}

	// fmt.Println(maxBottom)
	// fmt.Println(maxRight)

	res := 0
	for i := 0; i < count; i++ {
		for j := 0; j < len(treeMap[i]); j++ {
			val, _ := strconv.Atoi(fmt.Sprintf("%s", treeMap[i][j]))
			topIndexString := fmt.Sprintf("%d,%d", i-1, j)
			leftIndexString := fmt.Sprintf("%d,%d", i, j-1)
			bottomIndexString := fmt.Sprintf("%d,%d", i+1, j)
			rightIndexString := fmt.Sprintf("%d,%d", i, j+1)
			indexString := fmt.Sprintf("%d,%d", i, j)

			if i == 0 || j == 0 || i == len(treeMap)-1 || j == len(treeMap[i])-1 {
				fmt.Println(indexString)
				res += 1
			} else if maxTop[topIndexString] < val { // visible from top
				fmt.Println(indexString)
				res += 1
			} else if maxBottom[bottomIndexString] < val { // visible from bottom
				fmt.Println(indexString)
				res += 1
			} else if maxLeft[leftIndexString] < val { // visible from left
				fmt.Println(indexString)
				res += 1
			} else if maxRight[rightIndexString] < val { // visible from right
				fmt.Println(indexString)
				res += 1
			}
		}
	}

	fmt.Printf("part 1 is %d\n", res)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

	return
}

func part2() {
	maxScenicScore := 0
	// raytrace from each tree :")
	for i := 0; i < len(betterTreeMap); i++ {
		for j := 0; j < len(betterTreeMap[i]); j++ {
			if i == 0 || j == 0 || i == len(treeMap)-1 || j == len(treeMap[i])-1 {
				continue
			}

			score := 1
			// fmt.Printf("index is %d, %d\n", i, j)

			// up
			upScore := 0
			for m := i - 1; m >= 0; m-- {
				upScore += 1
				if betterTreeMap[m][j] >= betterTreeMap[i][j] {
					break
				}
			}
			score *= upScore

			// down
			downScore := 0
			for m := i + 1; m < len(treeMap); m++ {
				downScore += 1
				if betterTreeMap[m][j] >= betterTreeMap[i][j] {
					break
				}
			}
			score *= downScore

			// left
			leftScore := 0
			for m := j - 1; m >= 0; m-- {
				leftScore += 1
				if betterTreeMap[i][m] >= betterTreeMap[i][j] {
					break
				}
			}
			score *= leftScore

			// right
			rightScore := 0
			for m := j + 1; m < len(treeMap[i]); m++ {
				rightScore += 1
				if betterTreeMap[i][m] >= betterTreeMap[i][j] {
					break
				}
			}
			score *= rightScore

			// fmt.Printf("upscore is %d, downscore is %d, leftscore is %d, rightscore is %d\n", upScore, downScore, leftScore, rightScore)
			maxScenicScore = int(math.Max(float64(score), float64(maxScenicScore)))

		}
	}

	fmt.Printf("part 1 is %d\n", maxScenicScore)

}
