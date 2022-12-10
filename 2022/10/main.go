package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	cycles := map[int]int{20: 20, 60: 60, 100: 100, 140: 140, 180: 180, 220: 220}
	ccount := 0
	register := 1
	sum := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) == 1 {
			ccount += 1
			if val, ok := cycles[ccount]; ok {
				sum += (val * register)
				log.Println("1- ", val, register, val*register)
			}
		} else if len(line) == 2 {
			r, _ := strconv.Atoi(line[1])
			if val, ok := cycles[ccount+1]; ok {
				sum += (val * register)
				register += r
			} else {
				if val, ok := cycles[ccount+2]; ok {
					sum += (val * register)
				}
				register += r
			}
			ccount += 2
		}

	}

	log.Println(sum)
}
