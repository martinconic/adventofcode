package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func p1() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		card := strings.Split(line[1], "|")
		points := 0
		whand := strings.Fields(card[0])
		hand := strings.Fields(card[1])
		for _, c := range whand {
			if slices.Contains(hand, c) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points

	}

	fmt.Println(sum)

}

func p2() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	sum := 0
	insts := [219]int{}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		card := strings.Split(line[1], "|")
		points := 0
		whand := strings.Fields(card[0])
		hand := strings.Fields(card[1])
		for _, c := range whand {
			if slices.Contains(hand, c) {
				points++
			}
		}

		insts[i] += 1
		for j := 0; j < points; j++ {
			insts[i+j+1] += insts[i]
		}
		i++

	}
	for i := 0; i < len(insts); i++ {
		sum += insts[i]
	}
	fmt.Println(sum)

}

func main() {
	p1()
	p2()
}
