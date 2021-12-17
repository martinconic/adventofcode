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

		matrix = append(matrix, list)

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

		if i == l1-1 && j == l2-2 {
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
	min := 1000
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
