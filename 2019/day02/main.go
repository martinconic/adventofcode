package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	log.Println("Day2:")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	strvalues := strings.Split(scanner.Text(), ",")
	values := []int{}
	for _, v := range strvalues {
		intValue, _ := strconv.Atoi(v)
		values = append(values, intValue)
	}
	values[1] = 12
	values[2] = 2
	log.Println(values)

	for i := 0; i < len(values)-3; i += 4 {
		if values[i] == 99 {
			log.Println(values[0])
			return
		}
		if values[i] == 1 {
			values[values[i+3]] = values[values[i+1]] + values[values[i+2]]
		}
		if values[i] == 2 {
			values[values[i+3]] = values[values[i+1]] * values[values[i+2]]
		}
		log.Println(values)
	}
}

func partTwo() {
	log.Println("Day2:")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	strvalues := strings.Split(scanner.Text(), ",")
	ovalues := []int{}
	for _, v := range strvalues {
		intValue, _ := strconv.Atoi(v)
		ovalues = append(ovalues, intValue)
	}

	for x := 0; x <= 99; x++ {
		for y := 0; y <= 99; y++ {

			values := make([]int, len(ovalues))
			copy(values, ovalues)
			values[1] = x
			values[2] = y

			for i := 0; i < len(values)-3; i += 4 {
				// if values[i+3] >= len(values) || values[i+1] >= len(values) || values[i+2] >= len(values) {
				// break
				// }

				if values[i] == 99 {
					if values[0] == 19690720 {
						log.Println(values[0], 100*x+y)
						return
					}

					break
				}
				if values[i] == 1 {
					values[values[i+3]] = values[values[i+1]] + values[values[i+2]]
				}
				if values[i] == 2 {
					values[values[i+3]] = values[values[i+1]] * values[values[i+2]]
				}
			}

		}
	}
}

func main() {
	// partOne()
	partTwo()
}
