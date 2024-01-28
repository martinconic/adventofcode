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
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		records := strings.Fields(strings.ReplaceAll(line[0], ".", " "))
		arr := strings.Split(line[1], ",")

		fmt.Println(records, arr)

		count := 1
		if len(records) == len(arr) {
			for i, v := range records {
				av, _ := strconv.Atoi(arr[i])
				if len(v) > av {
					if !strings.Contains(string(v), "#") {
						count = len(v) * count
					}
				}
			}
		} else if len(records) < len(arr) {
			al, _ := strconv.Atoi(arr[len(arr)-1])
			if len(records[len(records)-1]) == al {
				records = records[:len(records)-1]
				arr = arr[:len(arr)-1]
			}
			a0, _ := strconv.Atoi(arr[0])
			if len(records[0]) == a0 {
				records = records[1:]
				arr = arr[1:]
			}
		}
		fmt.Println(count)
		sum += count
	}

	fmt.Println(sum)

}

func main() {
	p1()
}
