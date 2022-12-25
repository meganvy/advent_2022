// boilerplate
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int64
	throws      int64
	op          func(x int64) int64
	dividend    int64
	trueMonkey  int
	falseMonkey int
}

var monkeyMap map[int]*Monkey

func parseOp(operand, operator string) func(x int64) int64 {
	oldV, _ := strconv.Atoi(operator)
	v := int64(oldV)
	if operand == "+" {
		if operator == "old" {
			return func(x int64) int64 {
				return x + x
			}
		} else {
			return func(x int64) int64 {
				return x + v
			}
		}
	} else if operand == "*" {
		if operator == "old" {
			return func(x int64) int64 {
				return x * x
			}
		} else {
			return func(x int64) int64 {
				return x * v
			}
		}
	}
	return func(x int64) int64 {
		return x
	}

}

func main() {
	parseMonkeys()
}

func parseMonkeys() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input file")
		os.Exit(5)
	}

	monkeyMap := make(map[int]*Monkey)
	name := -1
	scanner := bufio.NewScanner(file)

	// parse inputs
	var monke int
	var operand string
	var operator string
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(line, "Monkey ") {
			name, err = strconv.Atoi(strings.Trim(line, "Monkey: "))
			monkey := Monkey{
				items:  []int64{},
				throws: int64(0),
			}
			m := &monkey
			monkeyMap[name] = m
		}
		if strings.HasPrefix(line, "Starting items:") {
			items := strings.Split(strings.Trim(line, " Starting items: "), ", ")
			intItems := make([]int64, len(items))
			for idx, i := range items {
				newI, _ := strconv.ParseInt(i, 10, 64)
				intItems[idx] = newI
			}
			monkeyMap[name].items = intItems
		}
		if strings.HasPrefix(line, "Test:") {
			dividend, _ := strconv.Atoi(strings.Trim(line, "Test: divisible by "))
			monkeyMap[name].dividend = int64(dividend)
		}
		if strings.HasPrefix(line, "Operation:") {
			fmt.Sscanf(line, "Operation: new = old %s %s", &operand, &operator)
			op := parseOp(operand, operator)
			monkeyMap[name].op = op
		}
		if strings.HasPrefix(line, "If true:") {
			fmt.Sscanf(line, "If true: throw to monkey %d", &monke)
			monkeyMap[name].trueMonkey = monke
		}
		if strings.HasPrefix(line, "If false:") {
			fmt.Sscanf(line, "If false: throw to monkey %d", &monke)
			if m, ok := monkeyMap[name]; ok {
				m.falseMonkey = monke
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(10)
	}

	// part1(monkeyMap)
	part2(monkeyMap)

	defer file.Close()
}

func part1(monkeyMap map[int]*Monkey) {
	fmt.Printf("%v", monkeyMap)
	// set up 20 rounds of monkey business
	for round := 0; round < 20; round++ {
		for m := 0; m < len(monkeyMap); m++ {
			monke := monkeyMap[m]
			for _, item := range monke.items {
				// monke inspect item
				monke.throws += 1
				newWorryLevel := monke.op(item)
				newWorryLevel = newWorryLevel / 3

				// monke test item
				if newWorryLevel%monke.dividend == 0 {
					monkeyMap[monke.trueMonkey].items = append(monkeyMap[monke.trueMonkey].items, newWorryLevel)
				} else {
					monkeyMap[monke.falseMonkey].items = append(monkeyMap[monke.falseMonkey].items, newWorryLevel)
				}
			}
			monke.items = []int64{}
		}
	}

	// calculate monkey business
	inspections := []int64{}
	for i := 0; i < len(monkeyMap); i++ {
		inspections = append(inspections, monkeyMap[i].throws)
	}
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	// fmt.Println(inspections[0] * inspections[1])

}

func part2(monkeyMap map[int]*Monkey) {
	// set up 10000 rounds of monkey business
	for round := 0; round < 1000; round++ {
		for m := 0; m < len(monkeyMap); m++ {
			monke := monkeyMap[m]
			for _, item := range monke.items {
				// monke inspect item
				newWorryLevel := monke.op(item)

				// monke test item
				if newWorryLevel%monke.dividend == 0 {
					monkeyMap[monke.trueMonkey].items = append(monkeyMap[monke.trueMonkey].items, newWorryLevel)
				} else {
					monkeyMap[monke.falseMonkey].items = append(monkeyMap[monke.falseMonkey].items, newWorryLevel)
				}
			}
			monke.throws += int64(len(monke.items))
			monke.items = []int64{}
		}
	}

	// calculate monkey business
	inspections := []int64{}
	for i := 0; i < len(monkeyMap); i++ {
		inspections = append(inspections, monkeyMap[i].throws)
	}
	fmt.Println(inspections)
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	fmt.Println(inspections[0] * inspections[1])

}
