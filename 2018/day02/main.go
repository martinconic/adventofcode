package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	twos := 0
	threes := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, "")
		m := map[string]int{}
		for _, v := range vals {
			if _, ok := m[v]; ok {
				m[v] += 1
			} else {
				m[v] = 1
			}
		}
		btwo := false
		bthree := false
		for _, v := range m {
			if v == 2 && !btwo {
				twos++
				btwo = true
			} else if v == 3 && !bthree {
				threes++
				bthree = true
			}
		}
	}

	log.Println(twos * threes)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	tvals := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, "")
		tvals = append(tvals, vals)
	}

	for i, v := range tvals {
		for j := i + 1; j < len(tvals); j++ {
			count := 0
			for k := 0; k < len(v); k++ {
				if v[k] == tvals[j][k] {
					count++
				}
			}
			if count == len(v)-1 {
				log.Println(v, tvals[j])
				return
			}

		}
	}

}

func main() {
	// partOne()
	partTwo()
}
