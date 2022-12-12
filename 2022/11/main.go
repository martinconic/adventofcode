package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	op1       string
	op2       string
	operation string
	divisible int
	iftrue    int
	iffalse   int
}

func worry_level(op1, op2 int, operand string) int {
	result := 0
	switch operand {
	case "+":
		result = int(math.Floor(float64(op1+op2) / 3.0))
	case "*":
		result = int(math.Floor(float64(op1*op2) / 3.0))
	default:
	}

	return result

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	monkeys := []Monkey{}

	m := map[int]int{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) > 0 {
			if line[0] == "Monkey" {
				monkey := Monkey{}
				monkeys = append(monkeys, monkey)
			}

			if line[0] == "Starting" {
				for i := 2; i < len(line); i++ {
					item, _ := strconv.Atoi(strings.ReplaceAll(line[i], ",", ""))
					monkeys[len(monkeys)-1].items = append(monkeys[len(monkeys)-1].items, item)
				}
			}

			if line[0] == "Operation:" {
				monkeys[len(monkeys)-1].op1 = line[3]
				monkeys[len(monkeys)-1].operation = line[4]
				monkeys[len(monkeys)-1].op2 = line[5]
			}

			if line[0] == "Test:" {
				monkeys[len(monkeys)-1].divisible, _ = strconv.Atoi(line[3])
			}

			if line[1] == "true:" {
				monkeys[len(monkeys)-1].iftrue, _ = strconv.Atoi(line[5])
			}
			if line[1] == "false:" {
				monkeys[len(monkeys)-1].iffalse, _ = strconv.Atoi(line[5])
			}
		}
	}
	for k := 0; k < 20; k++ {
		for i := 0; i < len(monkeys); i++ {
			m[i] += len(monkeys[i].items)
			for len(monkeys[i].items) > 0 {
				var o1, o2 int

				if monkeys[i].op1 == "old" {
					o1 = monkeys[i].items[0]
				} else {
					o1, _ = strconv.Atoi(monkeys[i].op1)
				}

				if monkeys[i].op2 == "old" {
					o2 = monkeys[i].items[0]
				} else {
					o2, _ = strconv.Atoi(monkeys[i].op2)
				}
				wl := worry_level(o1, o2, monkeys[i].operation)

				if wl%monkeys[i].divisible == 0 {
					monkeys[monkeys[i].iftrue].items = append(monkeys[monkeys[i].iftrue].items, wl)
				} else {
					monkeys[monkeys[i].iffalse].items = append(monkeys[monkeys[i].iffalse].items, wl)
				}

				if len(monkeys[i].items) < 2 {
					monkeys[i].items = []int{}
				} else {
					monkeys[i].items = monkeys[i].items[1:]
				}
			}
		}
	}

	log.Println(monkeys)
	log.Println(m)
}
