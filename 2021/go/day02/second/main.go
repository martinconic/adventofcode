package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	h := 0
	d := 0
	aim := 0
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(v[1])
		if v[0] == "forward" {
			d += value * aim
			h += value

			log.Println(h, d)
		} else if v[0] == "up" {
			// d -= value
			aim -= value
		} else if v[0] == "down" {
			// d += value
			aim += value
		}
	}

	log.Println(h, d, h*d)

}

//Answer:
