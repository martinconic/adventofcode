package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	counter := 0

	var depths []int

	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		deep, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, deep)
		l := len(depths)
		if l > 3 {
			if depths[l-1] > depths[l-4] {
				counter++
			}
		}
	}

	log.Println(counter)
}
