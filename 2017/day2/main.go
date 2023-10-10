package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		data := strings.Fields(scanner.Text())
		mx := 0
		mn := 100000
		for _, v := range data {
			value, _ := strconv.Atoi(v)
			if value > mx {
				mx = value
			}
			if value < mn {
				mn = value
			}
		}
		log.Println(mx, mn, mx-mn)
		sum += (mx - mn)
	}
	log.Println(sum)

}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		data := strings.Fields(scanner.Text())
		for i, sv := range data {
			value, _ := strconv.Atoi(sv)
			for j := i + 1; j < len(data); j++ {
				v, _ := strconv.Atoi(data[j])
				log.Println(value, v)
				if value > v {
					if value%v == 0 {
						sum += value / v
					}
				} else {
					if v%value == 0 {
						sum += v / value
					}
				}
			}
		}

	}
	log.Println(sum)

}

func main() {
	//	partOne()
	partTwo()
}
