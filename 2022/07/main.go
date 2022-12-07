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

	fs := []string{}
	m := map[string]int{}
	counter := 0

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[0] == "$" {
			if fields[1] == "cd" {
				if fields[2] != ".." {
					dir := fields[2]
					if _, ok := m[fields[2]]; ok {
						dir += string(counter)
						counter++
					}
					fs = append(fs, dir)

				} else {
					fs = fs[:len(fs)-1]
				}
			}
		} else {
			if fields[0] != "dir" {
				for i := 0; i < len(fs); i++ {
					size, _ := strconv.Atoi(fields[0])
					//log.Println(size, fs[i])
					m[fs[i]] += size
				}
			}
		}

	}

	//Part One

	//log.Println(m)
	sum := 0
	for _, v := range m {
		if v <= 100000 {
			sum += v
		}
	}

	log.Println(sum)

	//Part Two
	unused := 70000000 - m["/"]

	smallest := 30000000 - unused

	min := m["/"]

	for _, v := range m {
		if v >= smallest && v < min {
			min = v
		}
	}

	log.Println(min)
}
