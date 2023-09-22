package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	// "strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	sum := 0
	count := 0
	vals := []int{}
	m := map[int]int{}
	m[0] = 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		n, _ := strconv.Atoi(text)
		vals = append(vals, n)
		count++
		sum += n
		if _, ok := m[sum]; ok {
			log.Println(sum)
			return
		}
		m[sum] = 1
	}

	i := 0
	for {
		if i == len(vals) {
			i = 0
		}
		count++
		sum += vals[i]
		if _, ok := m[sum]; ok {
			log.Println(sum)
			return
		}
		m[sum] = 1
		i++
	}

}
