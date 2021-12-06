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
			h += value
		} else if v[0] == "up" {
			d -= value
		} else if v[0] == "down" {
			d += value
		}
	}

	log.Println(h)
	log.Println(d)
	log.Println(h * d)

}

//Answer: 1804520
