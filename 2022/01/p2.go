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

	max1 := 0
	max2 := 0
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
				max2 = max1
				max1 = max
				max = sum
			} else if sum > max1 {
				max2 = max1
				max1 = sum
			} else if sum > max2 {
				max2 = sum
			}

			sum = 0
		}
	}

	log.Println(max + max1 + max2)

}
