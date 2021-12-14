package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	elem := make(map[string]int)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := scanner.Text()

	rules := make(map[string]string)
	vmap := make(map[string]int)
	tempmap := make(map[string]int)

	scanner.Scan()
	for scanner.Scan() {
		str := scanner.Text()
		data := strings.Split(str, " -> ")
		rules[data[0]] = data[1]
	}

	// length := len(template)
	for i, letter := range template {
		elem[string(letter)] += 1
		if i > 0 {
			vmap[template[(i-1):(i+1)]] += 1
		}
	}

	for i := 0; i < 40; i++ {
		// log.Println(vmap)
		for key := range vmap {
			if val, ok := rules[key]; ok {
				// vmap[key] += 1
				elem[val] += vmap[key]
				key1 := string(key[0]) + val
				key2 := val + string(key[1])
				// log.Println(key, val, " ->", key1, key2)
				tempmap[key1] += vmap[key]
				tempmap[key2] += vmap[key]
			}
		}

		vmap = tempmap
		tempmap = make(map[string]int)
		// log.Println("vmap -> ", vmap)
	}

	max := 0
	min := math.MaxInt64

	log.Println("elem ->", elem)
	for _, value := range elem {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	log.Println(max - min)
}
