package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	pgames := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		impossible := false
		g := strings.Split(line, ":")

		sets := strings.Split(g[1], ";")

		for _, st := range sets {
			sg := strings.Split(st, ",")
			for _, ss := range sg {
				s := strings.Split(strings.TrimSpace(ss), " ")
				value, _ := strconv.Atoi(strings.TrimSpace(s[0]))
				//fmt.Println(s, s[0], s[1])
				if value > pgames[strings.TrimSpace(s[1])] {
					impossible = true
					break
				}
			}

			if impossible {
				break
			}
		}

		if impossible {
			continue
		}

		game, _ := strconv.Atoi(strings.Split(g[0], " ")[1])
		//fmt.Println(game)
		sum += game
	}
	fmt.Println(sum)

}

func p2() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		g := strings.Split(line, ":")

		sets := strings.Split(g[1], ";")
		mingame := map[string]int{"red": 1, "blue": 1, "green": 1}

		for _, st := range sets {
			sg := strings.Split(st, ",")
			for _, ss := range sg {
				s := strings.Split(strings.TrimSpace(ss), " ")
				value, _ := strconv.Atoi(s[0])
				if value > mingame[s[1]] {
					mingame[s[1]] = value
				}
				//fmt.Println(s, s[0], s[1])

			}

		}

		values := mingame["red"] * mingame["blue"] * mingame["green"]
		//		fmt.Println(values, mingame["red"], mingame["blue"], mingame["green"])
		sum += values
	}
	fmt.Println(sum)

}

func main() {
	p2()
}
