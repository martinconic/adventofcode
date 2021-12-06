package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	counter := 0
	previous := 0

	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Println(scanner.Text())
		deep, _ := strconv.Atoi(scanner.Text())
		if previous > 0 && deep > previous {
			counter++
		}
		previous = deep
	}

	log.Println(counter)
}

//Answer: 1681
