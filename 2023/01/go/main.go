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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		digits := []rune{}
		for _, l := range line {
			if unicode.IsDigit(l) {
				digits = append(digits, l)
			}
		}
		sval := string(digits[0]) + string(digits[len(digits)-1])
		val, _ := strconv.Atoi(sval)
		sum += val
	}

	fmt.Println(sum)

}

func p2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	sum := 0
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	scanner := bufio.NewScanner(file)
	index := 1
	for scanner.Scan() {
		first := 0
		last := 0
		ifirst := 1000
		ilast := -1
		line := scanner.Text()
		for k, d := range numbers {
			fi := strings.Index(line, d)
			if fi >= 0 {
				if fi < ifirst {
					first = k + 1
					ifirst = fi
				}
				if fi > ilast {
					last = k + 1
					ilast = fi
				}
			}

			li := strings.LastIndex(line, d)
			if li >= 0 {
				if li < ifirst {
					first = k + 1
					ifirst = li
				}
				if li > ilast {
					last = k + 1
					ilast = li
				}
			}

		}

		for k, d := range digits {
			fi := strings.Index(line, d)
			if fi >= 0 {
				if fi < ifirst {
					first = k + 1
					ifirst = fi
				}
				if fi > ilast {
					last = k + 1
					ilast = fi
				}
			}

			li := strings.LastIndex(line, d)
			if li >= 0 {
				if li < ifirst {
					first = k + 1
					ifirst = li
				}
				if li > ilast {
					last = k + 1
					ilast = li
				}
			}

		}
		fmt.Println(index, first, last)
		index++
		sum += first*10 + last
	}

	fmt.Println(sum)

}

func main() {
	//p1()
	p2()
}
