package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func p() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	count := 999999

	matrix := [][]string{}
	glxs := [][]int{}

	m := map[int]bool{}
	linm := map[int]bool{}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "")

		matrix = append(matrix, values)
		if !strings.Contains(line, "#") {
			linm[len(matrix)-1] = true
			continue
		}

		for i, v := range values {
			if v == "#" {
				glxs = append(glxs, []int{len(matrix) - 1, i})
				m[i] = true
			}
		}
	}

	ln := len(matrix)

	for i := ln - 1; i >= 0; i-- {
		if _, ok := linm[i]; ok {
			for j := 0; j < len(glxs); j++ {
				if glxs[j][0] > i {
					glxs[j][0] += count
				}
			}
		}
	}

	n := len(matrix[1])

	for i := n - 1; i >= 0; i-- {
		if _, ok := m[i]; !ok {
			for j := 0; j < len(glxs); j++ {
				if glxs[j][1] > i {
					glxs[j][1] += count
				}
			}
		}
	}

	sum := 0

	for i := 0; i < len(glxs); i++ {
		for j := i + 1; j < len(glxs); j++ {
			if glxs[i][0] < glxs[j][0] {
				sum += glxs[j][0] - glxs[i][0]
			} else {
				sum += glxs[i][0] - glxs[j][0]
			}
			if glxs[i][1] < glxs[j][1] {
				sum += glxs[j][1] - glxs[i][1]
			} else {
				sum += glxs[i][1] - glxs[j][1]
			}

		}
	}

	fmt.Println(sum)

}

func main() {
	p()
}
