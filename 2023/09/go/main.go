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
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hist := [][]int{}
		line := strings.Fields(scanner.Text())
		h := []int{}
		for _, v := range line {
			value, _ := strconv.Atoi(string(v))
			h = append(h, value)
		}

		hist = append(hist, h)

		pred := []int{h[len(h)-1]}

		j := 0
		count := 0
		for count != len(hist[j]) {
			h1 := []int{}
			count = 0
			for i := 1; i < len(hist[j]); i++ {
				value := hist[j][i] - hist[j][i-1]

				h1 = append(h1, value)
				if value == 0 {
					count++
				}

				if i == len(hist[j])-1 {
					pred = append(pred, value)
				}
			}
			hist = append(hist, h1)
			j++
		}
		ts := 0
		for i := len(pred) - 1; i >= 0; i-- {
			ts = pred[i] + ts
		}
		sum += ts
	}
	fmt.Println(sum)

}

func p2() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hist := [][]int{}
		line := strings.Fields(scanner.Text())
		h := []int{}
		for _, v := range line {
			value, _ := strconv.Atoi(string(v))
			h = append(h, value)
		}

		hist = append(hist, h)

		pred := []int{}

		j := 0
		count := 0
		for count != len(hist[j]) {
			h1 := []int{}
			count = 0
			for i := 1; i < len(hist[j]); i++ {
				value := hist[j][i] - hist[j][i-1]

				h1 = append(h1, value)
				if value == 0 {
					count++
				}

				if i == 1 {
					pred = append(pred, hist[j][i-1])
				}
			}
			hist = append(hist, h1)
			j++
		}
		ts := 0
		for i := 0; i < len(pred); i++ {
			if i == 0 || i%2 == 0 {
				ts = ts + pred[i]
			} else {
				ts = ts - pred[i]
			}
		}
		//	fmt.Println(pred)
		//	fmt.Println(ts)
		sum += ts
	}
	fmt.Println(sum)

}

func main() {
	p2()
}
