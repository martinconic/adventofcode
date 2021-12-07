package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var count int
var prevcount int

func main() {

	list := []int{}
	lmap := make(map[int]int)

	file, err := os.Open("../input.txt")

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

	for _, v := range list {
		lmap[v] += 1
	}

	log.Println(list)
	log.Println(len(list))
	log.Println(lmap)
	log.Println(len(lmap))

	max := list[len(list)-1]
	count := 0
	fuel := 0
	value := 0
	for i := 1; i <= max; i++ {
		for k, v := range lmap {
			if k > i {
				count += v * (k - i)
			} else {
				count += v * (i - k)
			}
			if count > fuel && fuel != 0 {
				break
			}
		}

		if count < fuel || fuel == 0 {
			fuel = count
			value = i
		}
		count = 0
	}

	log.Println("Position: ", value, "Fuel: ", fuel)

}
