package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	max := 0
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cal := scanner.Text()
		if len(cal) > 0 {
			val, _ := strconv.Atoi(cal)
			sum += val
		} else {
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}

	log.Println(max)

}
