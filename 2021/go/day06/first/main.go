package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	const days = 18

	list := []int{}
	file, err := os.Open("../sample.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	data := strings.Split(scanner.Text(), ",")

	for _, d := range data {
		n, _ := strconv.Atoi(d)
		list = append(list, n)
	}
	sort.Ints(list)

	temp := []int{}
	for i := 0; i < days; i++ {
		for j, value := range list {
			if value == 0 {
				list[j] = 6
				temp = append(temp, 8)
			} else {
				list[j] -= 1
			}
		}

		list = append(list, temp...)
		temp = []int{}
		log.Println(list)
	}

	log.Println(len(list))
}
