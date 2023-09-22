package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		fuel := value/3 - 2
		for i := fuel; i > 0; {
			sum += i
			i = i/3 - 2
		}
	}

	log.Println(sum)
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		sum += value/3 - 2
	}

	log.Println(sum)
}

func main() {
	// partOne()
	partTwo()
}
