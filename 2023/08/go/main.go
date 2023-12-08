package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func p1() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instr := strings.Split(scanner.Text(), "")
	scanner.Scan()

	nodes := map[string][]string{}
	el := "AAA"

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " = ")
		rl := strings.Split(line[1], ", ")

		values := []string{strings.ReplaceAll(rl[0], "(", ""), strings.ReplaceAll(rl[1], ")", "")}
		nodes[line[0]] = values

	}

	steps := 0
	i := 0
	for {
		steps++
		if instr[i] == "R" {
			el = nodes[el][1]
		} else {
			el = nodes[el][0]
		}
		//fmt.Println(el, nodes[el])

		if el == "ZZZ" {
			break
		}
		i++
		if len(instr) == i {
			i = 0
		}

	}

	fmt.Println(steps)
}

func p2() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instr := strings.Split(scanner.Text(), "")
	scanner.Scan()

	nodes := map[string][]string{}
	elements := []string{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " = ")
		if strings.LastIndex(line[0], "A") == len(line[0])-1 {
			elements = append(elements, line[0])
		}

		rl := strings.Split(line[1], ", ")
		values := []string{strings.ReplaceAll(rl[0], "(", ""), strings.ReplaceAll(rl[1], ")", "")}
		nodes[line[0]] = values

	}

	steps := 0
	i := 0
	m := map[string]int{}
	for {
		steps++
		if instr[i] == "R" {
			for j := 0; j < len(elements); j++ {
				node := nodes[elements[j]][1]
				elements[j] = node
				if strings.LastIndex(node, "Z") == len(node)-1 {
					if _, ok := m[node]; !ok {
						m[node] = steps
					}
				}
			}
		} else {
			for j := 0; j < len(elements); j++ {
				node := nodes[elements[j]][0]
				elements[j] = node
				if strings.LastIndex(node, "Z") == len(node)-1 {
					if _, ok := m[node]; !ok {
						m[node] = steps
					}

				}
			}

		}
		if len(m) == len(elements) {
			break
		}
		i++
		if len(instr) == i {
			i = 0
		}

	}

	res := []int{}
	for _, v := range m {
		res = append(res, v)
	}
	fmt.Println(lcm(res[0], res[1], res[2], res[3], res[4], res[5]))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
func main() {
	p2()
}
