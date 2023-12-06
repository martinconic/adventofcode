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
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	tim := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	scanner.Scan()
	dist := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	totalw := 1
	for i, t := range tim {
		tv, _ := strconv.Atoi(t)
		dv, _ := strconv.Atoi(dist[i])

		win := 0
		for k := 1; k < tv; k++ {
			score := k * (tv - k)
			if score > dv {
				win++
			}

			if score <= dv && win > 0 {
				break
			}
		}

		if win > 0 {
			totalw *= win
		}
	}

	fmt.Println(totalw)
}

func p2() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	t, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(scanner.Text(), ":")[1], " ", ""))

	scanner.Scan()
	dist, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(scanner.Text(), ":")[1], " ", ""))

	win := 0
	for k := 1; k < t; k++ {
		score := k * (t - k)
		if score > dist {
			win++
		}

		if score <= dist && win > 0 {
			break
		}
	}

	fmt.Println(win)
}

func main() {
	p2()
}
