package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int
var matrix [][]int
var l1, l2 int
var pq [][]int
var visited map[string]bool
var cost map[string]int

func main() {
	matrix = [][]int{}
	visited = make(map[string]bool)
	cost = make(map[string]int)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")
		list := []int{}
		for _, value := range data {
			val, _ := strconv.Atoi(value)
			list = append(list, val)
		}
		lngth := len(list)
		for i := 0; i < 4; i++ {
			for j := 0; j < lngth; j++ {
				val := list[j] + (i + 1)
				if val > 9 {
					val -= 9
				}

				list = append(list, val)
			}
		}

		matrix = append(matrix, list)
	}

	l1 = len(matrix)
	l2 = len(matrix[0])

	for c := 0; c < 4; c++ {
		for k := 0; k < l1; k++ {
			l := []int{}
			for i := 0; i < l2; i++ {
				val := matrix[k][i] + (c + 1)
				if val > 9 {
					val -= 9
				}
				l = append(l, val)
			}
			matrix = append(matrix, l)
		}
	}

	l1 = len(matrix)
	l2 = len(matrix[0])

	pq = [][]int{{0, 0, 0}}

	for {
		c, i, j := pop()

		if _, ok := visited[strconv.Itoa(i)+","+strconv.Itoa(j)]; ok {
			continue
		}
		visited[strconv.Itoa(i)+","+strconv.Itoa(j)] = true

		cost[strconv.Itoa(i)+","+strconv.Itoa(j)] = c

		if i == l1-1 && j == l2-1 {
			break
		}

		if i < l1-1 {
			pq = append(pq, []int{c + matrix[i+1][j], i + 1, j})
		}

		if i > 0 {
			pq = append(pq, []int{c + matrix[i-1][j], i - 1, j})
		}

		if j < l2-1 {
			pq = append(pq, []int{c + matrix[i][j+1], i, j + 1})
		}

		if j > 0 {
			pq = append(pq, []int{c + matrix[i][j-1], i, j - 1})
		}

	}

	log.Println(cost[strconv.Itoa(l1-1)+","+strconv.Itoa(l2-1)])
}

func pop() (c, i, j int) {
	min := 1000000
	index := 0
	for k, value := range pq {
		if min >= value[0] {
			min = value[0]
			c = value[0]
			i = value[1]
			j = value[2]
			index = k
		}
	}
	if index != len(pq) {
		pq = append(pq[:index], pq[index+1:]...)
	} else {
		pq = pq[:len(pq)-1]
	}

	return
}
