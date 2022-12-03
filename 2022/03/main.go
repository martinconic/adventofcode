package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	items := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		rucksack := scanner.Text()
		values := strings.Split(rucksack, "")
		h := len(values) / 2
		comp1 := strings.Join(values[0:h], "")
		comp2 := strings.Join(values[h:], "")

		seen := ""

		for _, c := range comp1 {
			if !strings.Contains(seen, string(c)) && strings.Contains(comp2, string(c)) {
				priority := strings.Index(items, string(c)) + 1
				log.Println(priority, string(c), comp1, comp2)
				sum += priority
				seen += string(c)
			}
		}
	}

	log.Println(sum)

}
