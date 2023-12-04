package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func p1() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	sum := 0

	numbers := map[string]string{}
	symbols := map[int][]int{}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	li := 0
	for scanner.Scan() {
		text := scanner.Text()
		isDigit := false
		snumber := ""
		k := ""
		for idx, r := range text {
			if unicode.IsDigit(r) {
				if !isDigit {
					k = strconv.Itoa(li) + "," + strconv.Itoa(idx)

					isDigit = true
				}
				snumber += string(r)
			} else {
				isDigit = false
				if len(k) > 0 {
					numbers[k] = snumber
					k = ""
					snumber = ""
				}
				if string(r) != "." {
					list := symbols[li]
					list = append(list, idx)
					symbols[li] = list
				}

			}

			if idx == len(text)-1 {
				if len(k) > 0 {
					numbers[k] = snumber
				}
			}
		}

		li++
	}

	//fmt.Println(numbers)
	//fmt.Println(symbols)

	for k, v := range numbers {
		keys := strings.Split(k, ",")
		y, _ := strconv.Atoi(keys[0])
		x1, _ := strconv.Atoi(keys[1])
		x2 := len(v) + x1
		found := false

		if values, ok := symbols[y]; ok {
			for _, value := range values {
				if value == x1-1 || value == x2 {
					val, _ := strconv.Atoi(v)
					sum += val
					found = true
					break
				}
			}
		}

		if found {
			continue
		}

		if values, ok := symbols[y-1]; ok {
			for _, value := range values {
				if value >= x1-1 && value <= x2 {
					val, _ := strconv.Atoi(v)
					sum += val
					found = true
					break
				}
			}
		}
		if found {
			continue
		}

		if values, ok := symbols[y+1]; ok {
			for _, value := range values {
				if value >= x1-1 && value <= x2 {
					val, _ := strconv.Atoi(v)
					sum += val
					break
				}
			}
		}

	}
	fmt.Println(sum)
}

func p2() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	sum := 0

	numbers := map[string]string{}
	symbols := map[int][]int{}
	adj := map[string][]int{}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	li := 0
	for scanner.Scan() {
		text := scanner.Text()
		isDigit := false
		snumber := ""
		k := ""
		for idx, r := range text {
			if unicode.IsDigit(r) {
				if !isDigit {
					k = strconv.Itoa(li) + "," + strconv.Itoa(idx)

					isDigit = true
				}
				snumber += string(r)
			} else {
				isDigit = false
				if len(k) > 0 {
					numbers[k] = snumber
					k = ""
					snumber = ""
				}
				if string(r) == "*" {
					list := symbols[li]
					list = append(list, idx)
					symbols[li] = list
				}

			}

			if idx == len(text)-1 {
				if len(k) > 0 {
					numbers[k] = snumber
				}
			}
		}

		li++
	}

	//fmt.Println(numbers)
	//fmt.Println(symbols)

	for k, v := range numbers {
		keys := strings.Split(k, ",")
		y, _ := strconv.Atoi(keys[0])
		x1, _ := strconv.Atoi(keys[1])
		x2 := len(v) + x1
		found := false

		if values, ok := symbols[y]; ok {
			for _, value := range values {
				if value == x1-1 || value == x2 {
					val, _ := strconv.Atoi(v)
					key := strconv.Itoa(y) + "," + strconv.Itoa(value)
					list := adj[key]
					list = append(list, val)
					adj[key] = list
					found = true
					break
				}
			}
		}

		if found {
			continue
		}

		if values, ok := symbols[y-1]; ok {
			for _, value := range values {
				if value >= x1-1 && value <= x2 {
					val, _ := strconv.Atoi(v)
					key := strconv.Itoa(y-1) + "," + strconv.Itoa(value)
					list := adj[key]
					list = append(list, val)
					adj[key] = list

					found = true
					break
				}
			}
		}
		if found {
			continue
		}

		if values, ok := symbols[y+1]; ok {
			for _, value := range values {
				if value >= x1-1 && value <= x2 {
					val, _ := strconv.Atoi(v)
					key := strconv.Itoa(y+1) + "," + strconv.Itoa(value)
					list := adj[key]
					list = append(list, val)
					adj[key] = list
					break
				}
			}
		}

	}
	//fmt.Println(adj)
	for _, v := range adj {
		if len(v) == 2 {
			value := v[0] * v[1]
			sum += value
		}
	}

	fmt.Println(sum)
}

func main() {
	//p1()
	p2()
}
