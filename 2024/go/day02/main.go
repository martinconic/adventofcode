package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		safe := true
		incr := true
		for i, _ := range line {
			if i > 0 {
				n1, _ := strconv.Atoi(line[i-1])
				n2, _ := strconv.Atoi(line[i])

				if i == 1 {
					if n1 > n2 {
						incr = false
					}
				}

				if (incr && (n1 >= n2 || n2 > n1+3)) || (!incr && (n1 <= n2 || n1 > n2+3)) {
					safe = false
					break
				}
			}
		}

		if safe {
			count++
		}
	}

	fmt.Println(count)

}

func p2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		for i := 0; i < len(line); i++ {
			l := make([]string, len(line))
			copy(l, line)
			l = append(l[:i], l[i+1:]...)
			// fmt.Println(l, line)
			if valid(l) {
				// fmt.Println(l, line)
				count++
				break
			}
		}

	}

	fmt.Println(count)

}

func valid(line []string) bool {
	safe := true
	incr := true
	n := len(line)

	for i := 0; i < n; i++ {
		if i > 0 {
			n1, _ := strconv.Atoi(line[i-1])
			n2, _ := strconv.Atoi(line[i])

			if i == 1 {
				if n1 > n2 {
					incr = false
				} else {
					incr = true
				}
			}

			// fmt.Println(line, tol, i, n1, n2, n)
			if (incr && (n1 >= n2 || n2 > n1+3)) || (!incr && (n1 <= n2 || n1 > n2+3)) {
				safe = false
				break
			}
		}
	}

	return safe

}

func main() {
	// p1()
	p2()
}
