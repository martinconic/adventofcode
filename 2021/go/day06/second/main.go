package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var count int
var prevcount int

func main() {

	const days = 256

	list := []int{}
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	data := strings.Split(scanner.Text(), ",")

	for _, d := range data {
		n, _ := strconv.Atoi(d)
		list = append(list, n)
	}

	sort.Ints(list)
	log.Println(list)
	previous := -1

	for _, d := range list {
		if d == previous {
			count += prevcount
			log.Println("Temp Count: ", count)
		} else {
			prevcount = 0
			previous = d

			step := days - (d + 1)
			count += 1
			prevcount += 1

			calculate(step)
			log.Println("Temp Count: ", count)
		}
	}
	log.Println(count)
}

func calculate(value int) {
	if value >= 9 {
		calculate(value - 9)
	}
	if value >= 7 {
		calculate(value - 7)
	}
	count += 1
	prevcount += 1
}
