// boilerplate
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	name           string
	subDirectories []*Directory
	size           int
	parent         *Directory
}

var sizes []int
var sizes2 []int
var res int64

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}

	sizes := []int{}
	sizes2 := []int{}
	scanner := bufio.NewScanner(file)
	dir := &Directory{
		name:           "/",
		subDirectories: []*Directory{},
		size:           0,
		parent:         nil,
	}
	curr := dir
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ ") {
			if line[2:4] == "cd" && line[5:] != "/" {
				// it's a directory node we've seen before
				p := false
				for _, subdir := range dir.subDirectories {
					if subdir.name == line[5:] {
						dir = subdir
						p = true
					}
				}
				// going up a directory
				if !p {
					if line[5:] == ".." {
						dir = dir.parent
					}
				}
			}
		} else if line[:3] == "dir" {
			tmp := dir
			newList := append(dir.subDirectories, &Directory{
				name:           line[4:],
				subDirectories: []*Directory{},
				size:           0,
				parent:         tmp,
			})
			dir.subDirectories = newList
		} else {
			fileLine := strings.Split(line, " ")
			size, _ := strconv.Atoi(fileLine[0])
			newList := append(dir.subDirectories, &Directory{
				name:           fileLine[1],
				subDirectories: []*Directory{},
				size:           size,
				parent:         dir,
			})
			dir.subDirectories = newList
		}
	}

	// part 1
	size(*curr, &sizes)
	// fmt.Printf("%v\n", sizes)

	sumSizes := 0
	for _, v := range sizes {
		if v <= 100000 {
			sumSizes += v
		}
	}
	fmt.Println(sumSizes)

	// part 2
	size2(*curr, &sizes2)
	// fmt.Println(sizes2)
	total := 0
	for _, size := range sizes2 {
		total = int(math.Max(float64(total), float64(size)))
	}
	unusedSpace := 70000000 - total
	sort.Slice(sizes2, func(i, j int) bool {
		return sizes2[i] < sizes2[j]
	})
	for _, val := range sizes2 {
		if unusedSpace+val >= 30000000 {
			fmt.Println(val)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

	defer file.Close()
}

func size(d Directory, sizes *[]int) int {
	subdirSize := 0
	if d.size != 0 {
		return d.size
	}
	if d.size == 0 {
		for _, s := range d.subDirectories {
			subdirSize += size(*s, sizes)
		}
		if subdirSize <= 100000 {
			*sizes = append(*sizes, subdirSize)
		}
	}
	return subdirSize
}

func size2(d Directory, sizes *[]int) int {
	subdirSize := 0
	if d.size != 0 {
		return d.size
	}
	if d.size == 0 {
		for _, s := range d.subDirectories {
			subdirSize += size2(*s, sizes)
		}
		*sizes = append(*sizes, subdirSize)
	}
	return subdirSize
}
