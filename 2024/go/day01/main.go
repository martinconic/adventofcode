package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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
	l1 := []int{}
	l2 := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		n1, _ := strconv.Atoi(fields[0])
		n2, _ := strconv.Atoi(fields[1])

		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	sort.Ints(l1)
	sort.Ints(l2)

	sum := 0

	for i := 0; i < len(l1); i++ {
		sum += diff(l1[i], l2[i])
	}

	log.Println(sum)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func p2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	l1 := []int{}
	l2 := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		n1, _ := strconv.Atoi(fields[0])

		l1 = append(l1, n1)
		l2 = append(l2, fields[1])
	}
	right := strings.Join(l2, ",")

	sum := 0

	for i := 0; i < len(l1); i++ {
		sum += l1[i] * strings.Count(right, strconv.Itoa(l1[i]))
	}

	log.Println(sum)
}

func main() {
	// p1()
	p2()
}
