package main

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Monkey64 struct {
	items     []*big.Int
	op1       string
	op2       string
	operation string
	divisible int
	iftrue    int
	iffalse   int
}

func worry_level2(op1, op2 *big.Int, operand string) *big.Int {
	result := big.NewInt(0)
	switch operand {
	case "+":
		result = result.Add(op1, op2)
	case "*":
		result = result.Mul(op1, op2)
	default:
	}

	return result

}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	monkeys := []Monkey64{}

	m := map[int]int{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) > 0 {
			if line[0] == "Monkey" {
				monkey := Monkey64{}
				monkeys = append(monkeys, monkey)
			}

			if line[0] == "Starting" {
				for i := 2; i < len(line); i++ {
					item, _ := strconv.Atoi(strings.ReplaceAll(line[i], ",", ""))
					bigitem := big.NewInt(int64(item))
					monkeys[len(monkeys)-1].items = append(monkeys[len(monkeys)-1].items, bigitem)
				}
			}

			if line[0] == "Operation:" {
				monkeys[len(monkeys)-1].op1 = line[3]
				monkeys[len(monkeys)-1].operation = line[4]
				monkeys[len(monkeys)-1].op2 = line[5]
			}

			if line[0] == "Test:" {
				div, _ := strconv.Atoi(line[3])
				monkeys[len(monkeys)-1].divisible = div
			}

			if line[1] == "true:" {
				ift, _ := strconv.Atoi(line[5])
				monkeys[len(monkeys)-1].iftrue = ift
			}
			if line[1] == "false:" {
				iff, _ := strconv.Atoi(line[5])
				monkeys[len(monkeys)-1].iffalse = iff
			}
		}
	}
	for k := 0; k < 1000; k++ {
		for i := 0; i < len(monkeys); i++ {
			m[i] += len(monkeys[i].items)
			for len(monkeys[i].items) > 0 {
				var o1, o2 *big.Int

				if monkeys[i].op1 == "old" {
					o1 = monkeys[i].items[0]

				} else {
					ot1, _ := strconv.Atoi(monkeys[i].op1)
					o1 = big.NewInt(int64(ot1))
				}

				if monkeys[i].op2 == "old" {
					o2 = monkeys[i].items[0]

				} else {
					ot2, _ := strconv.Atoi(monkeys[i].op2)
					o2 = big.NewInt(int64(ot2))
				}
				wl := worry_level2(o1, o2, monkeys[i].operation)

				div := big.NewInt(int64(monkeys[i].divisible))
				z := big.NewInt(1)
				z = z.Mod(wl, div)

				if z.Int64() == 0 {
					monkeys[monkeys[i].iftrue].items = append(monkeys[monkeys[i].iftrue].items, wl)
				} else {
					monkeys[monkeys[i].iffalse].items = append(monkeys[monkeys[i].iffalse].items, wl)
				}

				if len(monkeys[i].items) < 2 {
					monkeys[i].items = []*big.Int{}
				} else {
					monkeys[i].items = monkeys[i].items[1:]
				}
			}
		}
		log.Println(k)
	}

	//	log.Println(monkeys)
	log.Println(m)
}
